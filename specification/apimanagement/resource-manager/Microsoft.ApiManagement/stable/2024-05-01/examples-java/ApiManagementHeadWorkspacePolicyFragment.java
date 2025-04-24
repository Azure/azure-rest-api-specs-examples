
/**
 * Samples for WorkspacePolicyFragment GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspacePolicyFragment.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspacePolicyFragment.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspacePolicyFragment(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicyFragments().getEntityTagWithResponse("rg1", "apimService1", "wks1", "policyFragment1",
            com.azure.core.util.Context.NONE);
    }
}
