
/**
 * Samples for WorkspaceCertificate Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceCertificateWithKeyVault.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceCertificateWithKeyVault.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetWorkspaceCertificateWithKeyVault(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceCertificates().getWithResponse("rg1", "apimService1", "wks1", "templateCertkv",
            com.azure.core.util.Context.NONE);
    }
}
