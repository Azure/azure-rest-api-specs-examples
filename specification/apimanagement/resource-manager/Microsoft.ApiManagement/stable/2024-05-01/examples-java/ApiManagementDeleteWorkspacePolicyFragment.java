
/**
 * Samples for WorkspacePolicyFragment Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspacePolicyFragment.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspacePolicyFragment.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteWorkspacePolicyFragment(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicyFragments().deleteWithResponse("rg1", "apimService1", "wks1", "policyFragment1", "*",
            com.azure.core.util.Context.NONE);
    }
}
