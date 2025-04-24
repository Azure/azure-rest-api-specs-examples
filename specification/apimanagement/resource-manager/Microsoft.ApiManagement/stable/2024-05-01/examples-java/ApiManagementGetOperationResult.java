
/**
 * Samples for OperationsResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetOperationResult.json
     */
    /**
     * Sample code: ApiManagementGetOperationResult.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetOperationResult(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.operationsResults().getWithResponse("westus2", "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
            com.azure.core.util.Context.NONE);
    }
}
