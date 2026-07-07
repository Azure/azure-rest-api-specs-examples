
import com.azure.resourcemanager.storage.models.BlobInventoryPolicyName;

/**
 * Samples for BlobInventoryPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountGetBlobInventoryPolicy.json
     */
    /**
     * Sample code: StorageAccountGetBlobInventoryPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountGetBlobInventoryPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobInventoryPolicies().getWithResponse("res7687", "sto9699",
            BlobInventoryPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
