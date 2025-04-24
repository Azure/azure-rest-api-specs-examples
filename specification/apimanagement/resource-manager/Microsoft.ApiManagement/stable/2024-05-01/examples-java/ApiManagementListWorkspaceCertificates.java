
/**
 * Samples for WorkspaceCertificate ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceCertificates.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceCertificates.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceCertificates(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceCertificates().listByWorkspace("rg1", "apimService1", "wks1", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
