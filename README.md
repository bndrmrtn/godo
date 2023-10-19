![Thumbnail Image](https://github.com/bndrmrtn/godo/blob/main/thumbnails/thumbnail.jpg?raw=true)

## Installation

Firstly, you need to install the project dependencies:

```shell
go get
```

Then, build the project by running the following command:

```shell
go build
```

## Usage

To use GoDo, navigate to the project folder and execute the application from the command line by typing:

```shell
./godo
```

For a list of available options, you can run the help command:

```shell
./godo help
```

## How it works?

Here's an example of how GoDo works:

#### Create todo(s)
```
$ godo add "Clean my room" "Read a book"
✅  Todo added successfully
✅  Todo added successfully
Things to do:
⌚  1. Clean my room
⌚  2. Read a book
```

#### Finish todo(s)

```
$ godo finish 1
✅  Todo 1 set as done
Things to do:
✅  1. Clean my room
        • It took 14 minutes
⌚  2. Read a book
```

#### Other
For other examples download the project ;)