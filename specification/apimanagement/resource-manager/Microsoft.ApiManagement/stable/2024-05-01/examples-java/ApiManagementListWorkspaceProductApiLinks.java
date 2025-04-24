
/**
 * Samples for WorkspaceProductApiLink ListByProduct.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceProductApiLinks.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceProductApiLinks.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListWorkspaceProductApiLinks(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductApiLinks().listByProduct("rg1", "apimService1", "wks1", "product1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
