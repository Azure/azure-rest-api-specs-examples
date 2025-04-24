
/**
 * Samples for WorkspaceProductGroupLink ListByProduct.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceProductGroupLinks.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceProductGroupLinks.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListWorkspaceProductGroupLinks(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductGroupLinks().listByProduct("rg1", "apimService1", "wks1", "product1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
