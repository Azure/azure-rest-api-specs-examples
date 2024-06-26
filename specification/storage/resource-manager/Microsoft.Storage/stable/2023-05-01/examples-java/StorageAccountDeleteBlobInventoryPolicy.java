
import com.azure.resourcemanager.storage.models.BlobInventoryPolicyName;

/**
 * Samples for BlobInventoryPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * StorageAccountDeleteBlobInventoryPolicy.json
     */
    /**
     * Sample code: StorageAccountDeleteBlobInventoryPolicy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountDeleteBlobInventoryPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobInventoryPolicies().deleteWithResponse("res6977",
            "sto2527", BlobInventoryPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
