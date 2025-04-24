
/**
 * Samples for WorkspaceCertificate RefreshSecret.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementRefreshWorkspaceCertificate.json
     */
    /**
     * Sample code: ApiManagementRefreshWorkspaceCertificate.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementRefreshWorkspaceCertificate(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceCertificates().refreshSecretWithResponse("rg1", "apimService1", "wks1", "templateCertkv",
            com.azure.core.util.Context.NONE);
    }
}
