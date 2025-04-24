
/**
 * Samples for WorkspaceProduct GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceProduct.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceProduct.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceProduct(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProducts().getEntityTagWithResponse("rg1", "apimService1", "wks1", "unlimited",
            com.azure.core.util.Context.NONE);
    }
}
