# Overview
Pegasus is a library that allows very simple consuming of a kafka topic.  It utilizes the Shopify.sarama library to do so.

# Usage

This assumes a one-to-one relationship between a running service and a topic consumed.

Create a custom Processor that knows how to react to a message received.  This processor knows the topic it is going to watch, and the current offset.  It defaults to the newest offset.  It must be of the interface found in processor/processor.go

Instantiate a Pegasus instance, passing in a comma separated list of brokers.  Then pass the processor into the StartProcessor() function, and it will watch and react to messages.

# Example
```
pegasus := NewPegasus("kafka:9092")
exampleProcessor, _ := processor.NewExampleProcessor("example", sarama.OffsetNewest)
pegasus.StartProcessor(exampleProcessor)
```

# License and Author
Author: Tyler Ruppert

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.