
/**
 * Samples for WorkspaceDiagnostic GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceDiagnostic.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceDiagnostic.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceDiagnostic(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceDiagnostics().getEntityTagWithResponse("rg1", "apimService1", "wks1", "applicationinsights",
            com.azure.core.util.Context.NONE);
    }
}
