
import com.azure.resourcemanager.storagecache.models.NamespaceJunction;
import com.azure.resourcemanager.storagecache.models.Nfs3Target;
import com.azure.resourcemanager.storagecache.models.StorageTargetType;
import java.util.Arrays;

/**
 * Samples for StorageTargets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/
     * StorageTargets_CreateOrUpdate.json
     */
    /**
     * Sample code: StorageTargets_CreateOrUpdate.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void
        storageTargetsCreateOrUpdate(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargets().define("st1").withExistingCache("scgroup", "sc1")
            .withJunctions(Arrays.asList(
                new NamespaceJunction().withNamespacePath("/path/on/cache").withTargetPath("/path/on/exp1")
                    .withNfsExport("exp1").withNfsAccessPolicy("default"),
                new NamespaceJunction().withNamespacePath("/path2/on/cache").withTargetPath("/path2/on/exp2")
                    .withNfsExport("exp2").withNfsAccessPolicy("rootSquash")))
            .withTargetType(StorageTargetType.NFS3)
            .withNfs3(new Nfs3Target().withTarget("10.0.44.44").withUsageModel("READ_ONLY").withVerificationTimer(30))
            .create();
    }
}
