/** Samples for Certificate GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadCertificate.json
     */
    /**
     * Sample code: ApiManagementHeadCertificate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadCertificate(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .certificates()
            .getEntityTagWithResponse("rg1", "apimService1", "templateCert1", com.azure.core.util.Context.NONE);
    }
}
