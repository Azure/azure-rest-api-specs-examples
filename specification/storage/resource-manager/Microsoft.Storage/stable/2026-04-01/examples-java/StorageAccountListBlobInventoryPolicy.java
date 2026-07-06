
/**
 * Samples for BlobInventoryPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountListBlobInventoryPolicy.json
     */
    /**
     * Sample code: StorageAccountGetBlobInventoryPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountGetBlobInventoryPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobInventoryPolicies().list("res7687", "sto9699", com.azure.core.util.Context.NONE);
    }
}
