## What is an operator?
An operator is the most basic unit of log processing. Each operator fulfills a single responsibility, such as reading lines from a file, or parsing JSON from a field. Operators are then chained together in a pipeline to achieve a desired result.

For instance, a user may read lines from a file using the `file_input` operator. From there, the results of this operation may be sent to a `regex_parser` operator that creates fields based on a regex pattern. And then finally, these results may be sent to a `elastic_output` operator that writes each line to Elasticsearch.


## What operators are available?

Inputs:
- [file_input](./file_input.md)
- [journald_input](./journald_input.md)
- [k8s_event_input](./k8s_event_input.md)
- [stdin](./stdin.md)
- [syslog_input](./syslog_input.md)
- [tcp_input](./tcp_input.md)
- [udp_input](./udp_input.md)
- [windows_eventlog_input](./windows_eventlog_input.md)

Parsers:
- [csv_parser](./csv_parser.md)
- [json_parser](./json_parser.md)
- [regex_parser](./regex_parser.md)
- [syslog_parser](./syslog_parser.md)
- [severity_parser](./severity_parser.md)
- [time_parser](./time_parser.md)
- [trace_parser](./trace_parser.md)
- [uri_parser](./uri_parser.md)

Outputs:
- [file_output](./file_output.md)
- [stdout](./stdout.md)

General purpose:
- [add](./add.md)
- [copy](./copy.md)
- [filter](./filter.md)
- [flatten](./flatten.md)
- [move](./move.md)
- [recombine](./recombine.md)
- [remove](./remove.md)
- [retain](./retain.md)
- [router](./router.md)
