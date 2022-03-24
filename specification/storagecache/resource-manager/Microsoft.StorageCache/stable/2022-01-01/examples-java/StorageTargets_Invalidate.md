Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagecache_1.0.0-beta.5/sdk/storagecache/azure-resourcemanager-storagecache/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for StorageTargetOperation Invalidate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/StorageTargets_Invalidate.json
     */
    /**
     * Sample code: StorageTargets_Invalidate.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsInvalidate(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargetOperations().invalidate("scgroup", "sc", "st1", Context.NONE);
    }
}
```
