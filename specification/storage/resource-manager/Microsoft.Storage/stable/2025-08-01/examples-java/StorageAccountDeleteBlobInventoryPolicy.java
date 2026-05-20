
import com.azure.resourcemanager.storage.models.BlobInventoryPolicyName;

/**
 * Samples for BlobInventoryPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountDeleteBlobInventoryPolicy.json
     */
    /**
     * Sample code: StorageAccountDeleteBlobInventoryPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountDeleteBlobInventoryPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobInventoryPolicies().deleteWithResponse("res6977", "sto2527",
            BlobInventoryPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
