
/**
 * Samples for WorkspaceProduct Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceProduct.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceProduct.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceProduct(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProducts().deleteWithResponse("rg1", "apimService1", "wks1", "testproduct", "*", true,
            com.azure.core.util.Context.NONE);
    }
}
