import com.azure.core.util.Context;

/** Samples for Certificate RefreshSecret. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementRefreshCertificate.json
     */
    /**
     * Sample code: ApiManagementRefreshCertificate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementRefreshCertificate(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.certificates().refreshSecretWithResponse("rg1", "apimService1", "templateCertkv", Context.NONE);
    }
}
