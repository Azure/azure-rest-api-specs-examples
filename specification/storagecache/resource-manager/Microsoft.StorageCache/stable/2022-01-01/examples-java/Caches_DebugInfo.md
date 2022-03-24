Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagecache_1.0.0-beta.5/sdk/storagecache/azure-resourcemanager-storagecache/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Caches DebugInfo. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Caches_DebugInfo.json
     */
    /**
     * Sample code: Caches_DebugInfo.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesDebugInfo(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().debugInfo("scgroup", "sc", Context.NONE);
    }
}
```
