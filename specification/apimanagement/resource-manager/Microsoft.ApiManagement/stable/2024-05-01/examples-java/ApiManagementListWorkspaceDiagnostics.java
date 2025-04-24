
/**
 * Samples for WorkspaceDiagnostic ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceDiagnostics.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceDiagnostics.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceDiagnostics(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceDiagnostics().listByWorkspace("rg1", "apimService1", "wks1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
