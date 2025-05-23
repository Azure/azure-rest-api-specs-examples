
/**
 * Samples for Certificate Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteCertificate.json
     */
    /**
     * Sample code: ApiManagementDeleteCertificate.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteCertificate(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.certificates().deleteWithResponse("rg1", "apimService1", "tempcert", "*",
            com.azure.core.util.Context.NONE);
    }
}
