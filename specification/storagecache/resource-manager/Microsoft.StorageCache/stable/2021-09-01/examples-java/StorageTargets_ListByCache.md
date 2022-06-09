```java
import com.azure.core.util.Context;

/** Samples for StorageTargets ListByCache. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/StorageTargets_ListByCache.json
     */
    /**
     * Sample code: StorageTargets_List.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargets().listByCache("scgroup", "sc1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagecache_1.0.0-beta.4/sdk/storagecache/azure-resourcemanager-storagecache/README.md) on how to add the SDK to your project and authenticate.
