
/**
 * Samples for WorkspaceSubscription Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceSubscription.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceSubscription.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceSubscription(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceSubscriptions().deleteWithResponse("rg1", "apimService1", "wks1", "testsub", "*",
            com.azure.core.util.Context.NONE);
    }
}
