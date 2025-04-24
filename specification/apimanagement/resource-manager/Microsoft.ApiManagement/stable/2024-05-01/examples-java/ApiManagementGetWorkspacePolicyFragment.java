
/**
 * Samples for WorkspacePolicyFragment Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspacePolicyFragment.json
     */
    /**
     * Sample code: ApiManagementGetWorkspacePolicyFragment.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspacePolicyFragment(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicyFragments().getWithResponse("rg1", "apimService1", "wks1", "policyFragment1", null,
            com.azure.core.util.Context.NONE);
    }
}
