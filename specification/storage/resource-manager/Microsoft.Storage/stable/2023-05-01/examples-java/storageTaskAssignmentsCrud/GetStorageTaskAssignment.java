
/**
 * Samples for StorageTaskAssignments Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/storageTaskAssignmentsCrud/
     * GetStorageTaskAssignment.json
     */
    /**
     * Sample code: GetStorageTaskAssignment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getStorageTaskAssignment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageTaskAssignments().getWithResponse("res4228",
            "sto4445", "myassignment1", com.azure.core.util.Context.NONE);
    }
}
