
/**
 * Samples for WorkspaceDiagnostic Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceDiagnostic.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceDiagnostic.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceDiagnostic(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceDiagnostics().getWithResponse("rg1", "apimService1", "wks1", "applicationinsights",
            com.azure.core.util.Context.NONE);
    }
}
