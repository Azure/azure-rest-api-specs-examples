
/**
 * Samples for WorkspacePolicy ListByApi.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspacePolicies.json
     */
    /**
     * Sample code: ApiManagementListWorkspacePolicies.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspacePolicies(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicies().listByApi("rg1", "apimService1", "wks1", com.azure.core.util.Context.NONE);
    }
}
