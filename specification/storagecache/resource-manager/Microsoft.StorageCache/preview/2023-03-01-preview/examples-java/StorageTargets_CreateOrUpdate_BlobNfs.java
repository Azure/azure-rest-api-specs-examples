import com.azure.resourcemanager.storagecache.models.BlobNfsTarget;
import com.azure.resourcemanager.storagecache.models.NamespaceJunction;
import com.azure.resourcemanager.storagecache.models.StorageTargetType;
import java.util.Arrays;

/** Samples for StorageTargets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-03-01-preview/examples/StorageTargets_CreateOrUpdate_BlobNfs.json
     */
    /**
     * Sample code: StorageTargets_CreateOrUpdate_BlobNfs.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsCreateOrUpdateBlobNfs(
        com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager
            .storageTargets()
            .define("st1")
            .withExistingCache("scgroup", "sc1")
            .withJunctions(Arrays.asList(new NamespaceJunction().withNamespacePath("/blobnfs")))
            .withTargetType(StorageTargetType.BLOB_NFS)
            .withBlobNfs(
                new BlobNfsTarget()
                    .withTarget(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/blofnfs/blobServices/default/containers/blobnfs")
                    .withUsageModel("READ_WRITE")
                    .withVerificationTimer(28800)
                    .withWriteBackTimer(3600))
            .create();
    }
}
