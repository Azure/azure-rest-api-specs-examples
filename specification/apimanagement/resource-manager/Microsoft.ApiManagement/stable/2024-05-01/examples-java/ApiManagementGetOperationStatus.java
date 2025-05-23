
/**
 * Samples for OperationStatus Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetOperationStatus.json
     */
    /**
     * Sample code: Get operation status.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void getOperationStatus(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.operationStatus().getWithResponse("testLocation", "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
            com.azure.core.util.Context.NONE);
    }
}
