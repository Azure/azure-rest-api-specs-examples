Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagecache_1.0.0-beta.4/sdk/storagecache/azure-resourcemanager-storagecache/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.storagecache.models.Nfs3Target;
import com.azure.resourcemanager.storagecache.models.StorageTargetType;

/** Samples for StorageTargets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/StorageTargets_CreateOrUpdate_NoJunctions.json
     */
    /**
     * Sample code: StorageTargets_CreateOrUpdate_NoJunctions.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsCreateOrUpdateNoJunctions(
        com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager
            .storageTargets()
            .define("st1")
            .withExistingCache("scgroup", "sc1")
            .withTargetType(StorageTargetType.NFS3)
            .withNfs3(new Nfs3Target().withTarget("10.0.44.44").withUsageModel("READ_HEAVY_INFREQ"))
            .create();
    }
}
```
