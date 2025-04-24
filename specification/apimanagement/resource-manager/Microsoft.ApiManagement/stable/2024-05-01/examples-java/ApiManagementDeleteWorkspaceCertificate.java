
/**
 * Samples for WorkspaceCertificate Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceCertificate.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceCertificate.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceCertificate(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceCertificates().deleteWithResponse("rg1", "apimService1", "wks1", "tempcert", "*",
            com.azure.core.util.Context.NONE);
    }
}
