/** Samples for Certificate RefreshSecret. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementRefreshCertificate.json
     */
    /**
     * Sample code: ApiManagementRefreshCertificate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementRefreshCertificate(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .certificates()
            .refreshSecretWithResponse("rg1", "apimService1", "templateCertkv", com.azure.core.util.Context.NONE);
    }
}
