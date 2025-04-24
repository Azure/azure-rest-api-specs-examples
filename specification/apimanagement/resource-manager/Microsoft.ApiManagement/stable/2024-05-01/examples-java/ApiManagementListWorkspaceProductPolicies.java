
/**
 * Samples for WorkspaceProductPolicy ListByProduct.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceProductPolicies.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceProductPolicies.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListWorkspaceProductPolicies(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductPolicies().listByProductWithResponse("rg1", "apimService1", "wks1",
            "armTemplateProduct4", com.azure.core.util.Context.NONE);
    }
}
