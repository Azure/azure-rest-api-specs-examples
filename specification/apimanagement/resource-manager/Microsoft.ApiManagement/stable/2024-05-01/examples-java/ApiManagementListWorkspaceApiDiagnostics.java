
/**
 * Samples for WorkspaceApiDiagnostic ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceApiDiagnostics.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceApiDiagnostics.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceApiDiagnostics(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiDiagnostics().listByWorkspace("rg1", "apimService1", "wks1", "echo-api", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
