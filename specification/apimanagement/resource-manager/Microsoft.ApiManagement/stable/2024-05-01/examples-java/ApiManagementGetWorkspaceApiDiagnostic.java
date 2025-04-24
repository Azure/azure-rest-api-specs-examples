
/**
 * Samples for WorkspaceApiDiagnostic Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceApiDiagnostic.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceApiDiagnostic.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceApiDiagnostic(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiDiagnostics().getWithResponse("rg1", "apimService1", "wks1", "57d1f7558aa04f15146d9d8a",
            "applicationinsights", com.azure.core.util.Context.NONE);
    }
}
