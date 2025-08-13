# Create a cache server of web pages with REST API. 

### Functional requirements:
*Operations*
1) "put" operation - association of key (page URL) with page bytes + increases page hit rating.
2) "get" operation - fetch of page bytes by key (page URL) + increases page hit rating.
3) "delete" operation - removes a key from the cache. 
4) "top" operation - returns ordered list of top 100 keys sorted by hit rating desc.
5) items that were not used during some configurable amount of time should be evicted from the cache (TTL).

*Configs*
1) Config param that limits max count for stored items.
2) Config param that limits the max size of a single item in bytes.
3) Config param that limits the max size of the whole cache in bytes.


### Non-functional requirements:
1) The cache should support concurrent access.
2) All operations should not have linear or worse complexity dependence on the number of stored items.
3) If the cache is full in terms of max count or max overall size limits, it should evict the oldest item (among the equally old elements it should evict one with a lower hit rating).
