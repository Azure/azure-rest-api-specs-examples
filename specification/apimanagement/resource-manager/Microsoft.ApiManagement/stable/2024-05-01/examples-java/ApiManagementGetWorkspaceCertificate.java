
/**
 * Samples for WorkspaceCertificate Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceCertificate.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceCertificate.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceCertificate(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceCertificates().getWithResponse("rg1", "apimService1", "wks1", "templateCert1",
            com.azure.core.util.Context.NONE);
    }
}
