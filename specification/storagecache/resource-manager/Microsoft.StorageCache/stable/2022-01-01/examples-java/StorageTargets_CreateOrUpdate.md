Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagecache_1.0.0-beta.5/sdk/storagecache/azure-resourcemanager-storagecache/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.storagecache.models.NamespaceJunction;
import com.azure.resourcemanager.storagecache.models.Nfs3Target;
import com.azure.resourcemanager.storagecache.models.StorageTargetType;
import java.util.Arrays;

/** Samples for StorageTargets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/StorageTargets_CreateOrUpdate.json
     */
    /**
     * Sample code: StorageTargets_CreateOrUpdate.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsCreateOrUpdate(
        com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager
            .storageTargets()
            .define("st1")
            .withExistingCache("scgroup", "sc1")
            .withJunctions(
                Arrays
                    .asList(
                        new NamespaceJunction()
                            .withNamespacePath("/path/on/cache")
                            .withTargetPath("/path/on/exp1")
                            .withNfsExport("exp1")
                            .withNfsAccessPolicy("default"),
                        new NamespaceJunction()
                            .withNamespacePath("/path2/on/cache")
                            .withTargetPath("/path2/on/exp2")
                            .withNfsExport("exp2")
                            .withNfsAccessPolicy("rootSquash")))
            .withTargetType(StorageTargetType.NFS3)
            .withNfs3(new Nfs3Target().withTarget("10.0.44.44").withUsageModel("READ_HEAVY_INFREQ"))
            .create();
    }
}
```
