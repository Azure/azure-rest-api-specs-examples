
/**
 * Samples for WorkspaceDiagnostic Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceDiagnostic.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceDiagnostic.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceDiagnostic(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceDiagnostics().deleteWithResponse("rg1", "apimService1", "wks1", "applicationinsights", "*",
            com.azure.core.util.Context.NONE);
    }
}
