/** Samples for Certificate Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetCertificate.json
     */
    /**
     * Sample code: ApiManagementGetCertificate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetCertificate(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .certificates()
            .getWithResponse("rg1", "apimService1", "templateCert1", com.azure.core.util.Context.NONE);
    }
}
