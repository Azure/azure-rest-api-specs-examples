
/**
 * Samples for WorkspaceCertificate GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceCertificate.json
     */
    /**
     * Sample code: ApiManagementWorkspaceHeadCertificate.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementWorkspaceHeadCertificate(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceCertificates().getEntityTagWithResponse("rg1", "apimService1", "wks1", "templateCert1",
            com.azure.core.util.Context.NONE);
    }
}
