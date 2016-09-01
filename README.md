# Overview
Pegasus is a service that consumes a kafka topic and acts on different keys passed to it.

# Usage

# Example
```
docker run --link kafka:kafka -e DOCKERCLOUD_USER=[DOCKERCLOUD_USER_HERE] -e DOCKERCLOUD_APIKEY=[DOCKERCLOUD_APIKEY_HERE] ruppdog/pegasus
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