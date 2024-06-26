
/**
 * Samples for AppServiceCertificateOrders DeleteCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/
     * DeleteAppServiceCertificate.json
     */
    /**
     * Sample code: Delete App Service Certificate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAppServiceCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().deleteCertificateWithResponse(
            "testrg123", "SampleCertificateOrderName", "SampleCertName1", com.azure.core.util.Context.NONE);
    }
}
