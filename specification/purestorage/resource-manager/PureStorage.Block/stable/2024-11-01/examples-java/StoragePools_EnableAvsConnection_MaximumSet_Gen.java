
import com.azure.resourcemanager.purestorageblock.models.StoragePoolEnableAvsConnectionPost;

/**
 * Samples for StoragePools EnableAvsConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/StoragePools_EnableAvsConnection_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_EnableAvsConnection.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        storagePoolsEnableAvsConnection(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().enableAvsConnection("rgpurestorage", "storagePoolname",
            new StoragePoolEnableAvsConnectionPost().withClusterResourceId("tghkgktlddwlszbeh"),
            com.azure.core.util.Context.NONE);
    }
}
