
/**
 * Samples for WorkspaceApiDiagnostic Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceApiDiagnostic.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceApiDiagnostic.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteWorkspaceApiDiagnostic(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiDiagnostics().deleteWithResponse("rg1", "apimService1", "wks1", "57d1f7558aa04f15146d9d8a",
            "applicationinsights", "*", com.azure.core.util.Context.NONE);
    }
}
