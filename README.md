# rocket-messages

This is a small service which accepts rocket event messages, stores them and
calculates the current state of each of the rockets.

## Running the code

The project uses [Taskfile](https://taskfile.dev) as task runner. [Installation
instructions](https://taskfile.dev/installation/). This is mainly due to it
being more modern than eg. `make`.

See the available tasks using `task -l`. Some tasks depend on a few tools being
installed. The requirements are listed here:

- `go`
- `watch`
- `jq`

Apart from these the `ROCKETS_BINARY` environment variable can be used to point
to what binary should be used for starting the service which sends messages.
Alternatively you can place the binary in the `./bin` folder.

- Use `task run` to start listening for rockets.
- Use `task run-rockets` to start sending messages.
- Use `task watch-rockets` to watch the current state of the rockets.

The tasks are blocking and should each be run in their own shell.

## Design

As the input is messages, it made sense to do a simple version of event
sourcing, in order to demonstrate understading of the concept.

Incoming messages are stored in a `map[string][]Message` indexed by the channel
ID. This makes it easy to fetch all messages for a given rocket. When getting a
rocket or listing them all, the state is calculated based on the stored
messages.

Currently the design doesn't have a concept of snapshotting, which means that
performance will suffer as the amount of messages grows.

The store is abstracted by an interface so that one can easily implement a
database backed store. For the purposes of this assignment, I deemed in-memory
to be sufficient.

### Data gurantees

Messages are stored ordered by their message number. Currently this is naively
done by sorting the messages on each insert. Even if the message history has a
gap, the most recent message is included in the state calculation. And if an
older message arrives later then it will be inserted in the right place in the
sequence, and new calculations will include that message.

I considered to implement a buffer for storing out-of-order messages, but did
it like this instead for simplicity.

### Testing

I wrote a few tests of the message store and the rocket service. These can be
run using `task test`.

## Final thoughts

The code was written tiredly in the very short breaks in a 3 hour infant
feeding cycle, and might contain traces of this 😅 We got a pair of twins a
week ago, and they need a bit more care in the first few weeks.

I didn't prioritize implementing the sorting of the rocket list, as I deemed it
less interesting compared to the other parts of the project.
