use std::fs;
use std::env;
use std::io::Write;

use clap::{Parser, Subcommand};
use chrono::{Duration, NaiveDateTime, Utc};
use serde::{Deserialize, Serialize};

fn main() {
    let cli = Cli::parse();
    match &cli.command {
        Commands::Init { work, rest } => {
            init_timer(*work, *rest);
        },
        _ => {
            print_remaining_time(read_state());
        }
    }
}

#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
#[clap(propagate_version = true)]
struct Cli {
    #[clap(subcommand)]
    command: Commands,
}

/// Extremely simple pomodoro timer
#[derive(Subcommand, Debug)]
enum Commands {
    /// Initializes a timer with specified durations
    Init {
        /// Duration of work timer (minutes)
        #[clap(short, long, default_value_t = 25)]
        work: i64,
        /// Duration of break timer (minutes)
        #[clap(short, long, default_value_t = 5)]
        rest: i64,
    },
    /// Prints remaining time on timer to stdout
    Show
}

#[derive(Debug, Clone, Serialize, Deserialize)]
struct State {
    timestamp: i64,
    work_dur: i64,
    rest_dur: i64,
    is_break: bool,
}

/// Wrapper function for state instantiation, duration is in minutes.
fn init_timer(work_dur: i64, rest_dur: i64) {
    let state = State {
        timestamp: Utc::now().timestamp(),
        work_dur: work_dur * 60, // convert to seconds for unix time
        rest_dur: rest_dur * 60,
        is_break: false, // timer just started... you don't get a break yet!
    };
    write_file(state);
}

/// writes state to /tmp file "database". writing to this file is essentially
/// resetting the timer state.
fn write_file(state: State) {
    let json = serde_json::to_string(&state).unwrap();
    let db = env::temp_dir().join("pomo.doro");
    let mut f = fs::File::create(db).expect("error opening file");
    f.write(json.as_bytes()).expect("error writing file");
}

/// Reads the temporary file database and deserializes
fn read_state() -> State {
    let db = env::temp_dir().join("pomo.doro");
    let contents = fs::read_to_string(db).expect("error reading file");
    serde_json::from_str(&contents).expect("error parsing state")
}

/// Prints to the user the remaining time
fn print_remaining_time(state: State) {
    let stamp = NaiveDateTime::from_timestamp(state.timestamp, 0);
    let now = Utc::now().naive_utc();
    let duration = now.signed_duration_since(stamp);
    let remaining = if state.is_break {
        Duration::seconds(state.rest_dur - duration.num_seconds())
    } else {
        Duration::seconds(state.work_dur - duration.num_seconds())
    };


    // reset our timer, swap to other mode (ie work <-> break)
    if remaining.num_seconds() <= 0 {
        write_file(State {
            timestamp: Utc::now().timestamp(),
            work_dur: state.work_dur,
            rest_dur: state.rest_dur,
            is_break: !state.is_break,
        });
        println!("timer complete!");
    } else if state.is_break {
        println!("{}: {:0>2}:{:0>2}", "break", remaining.num_minutes(), remaining.num_seconds() % 60);
    } else {
        println!("{}: {:0>2}:{:0>2}", "work", remaining.num_minutes(), remaining.num_seconds() % 60);
    }
}
