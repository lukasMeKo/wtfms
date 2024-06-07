# What this?

A PoC showing how to store manifold articles in a document database.
Other WMS functionality is intentionally left out, since the author believes,
that those processes are better mapped to a SQL database.

We store JSON schemas for properties and build categories as (ordered) sets of properties.
An article index is then created for each category, where the index is derived from the category schema.
An article is an object with properties, where each property must be present in the category and be valid given the properties schema.
Additional validation can be performed by registering a property handler in the go code.

# How to test?

```
make
```

You will see output, note that we correctly caught malformed articles, both by schema and by validator!
Your browser might open and you should see the articles in the meilisearch-ui.

# References

- <https://www.meilisearch.com/>: elastic-search like engine, easier to setup and probably more resource efficent and less featurefull.
- <https://json-schema.org/>: Specify allowed shape of an json object
- <https://github.com/rjsf-team/react-jsonschema-form>: generate components from json schema
- <https://github.com/spf13/cobra>: cli lib for go
