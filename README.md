# CLI Task Manager

CLI Task Manager is a command-line interface tool for managing your TODOs directly from your terminal.

## Overview

The basic usage of the tool involves adding new tasks, listing incomplete tasks, and marking tasks as complete. Below is a quick overview of available commands:

- `add`: Add a new task to your TODO list.
- `list`: List all of your incomplete tasks.
- `do`: Mark a task on your TODO list as complete.

## Installation

### Prerequisites

- Go programming language installed on your machine.

### Instructions

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Acs176/cli-task-manager.git
   ```
2. Navigate to the directory and install the CLI Task Manager
   ```bash
   cd cli-task-manager
   go install
   ```

## Usage
After installation, you can start using the CLI Task Manager by typing cli-task-manager followed by a command. Below are some examples of usage:
  ```bash
  # Display available commands
  cli-task-manager
  
  # Add a new task
  cli-task-manager add "review talk proposal"
  
  # List all incomplete tasks
  cli-task-manager list
  
  # Mark a task as complete
  cli-task-manager do 1
  ```
  You can optionally change the name of the installed binary to make the keyword shorter
## Development

### Libraries Used

- [spf13/cobra](https://github.com/spf13/cobra): Used for building the CLI shell.

### Database

- BoltDB is used for interacting with the database. You can find more information about BoltDB [here](https://github.com/boltdb/bolt).

