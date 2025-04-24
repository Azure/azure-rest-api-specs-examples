
/**
 * Samples for WorkspaceProduct Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceProduct.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceProduct.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceProduct(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProducts().getWithResponse("rg1", "apimService1", "wks1", "unlimited",
            com.azure.core.util.Context.NONE);
    }
}
