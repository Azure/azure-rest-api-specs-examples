
/**
 * Samples for WorkspacePolicyFragment ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspacePolicyFragments.json
     */
    /**
     * Sample code: ApiManagementListWorkspacePolicyFragments.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListWorkspacePolicyFragments(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicyFragments().listByService("rg1", "apimService1", "wks1", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
