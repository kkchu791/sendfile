### results

From running this I found the zero copy approach is 11x faster than the copy approach.

Kafka uses zero copy by default for faster speeds. But if you want the data encrypted, it would have to use the copy approach so you can encrypt the data in the user space.

#### benchmark output

```
The copy way took 155.834µs to run.

The zero copy way took 27.25µs to run.
```
