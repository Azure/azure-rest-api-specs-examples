
/**
 * Samples for WorkspaceProduct ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceProducts.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceProducts.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceProducts(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProducts().listByService("rg1", "apimService1", "wks1", null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
