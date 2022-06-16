import com.azure.core.util.Context;

/** Samples for Certificate Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetCertificateWithKeyVault.json
     */
    /**
     * Sample code: ApiManagementGetCertificateWithKeyVault.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetCertificateWithKeyVault(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.certificates().getWithResponse("rg1", "apimService1", "templateCertkv", Context.NONE);
    }
}
