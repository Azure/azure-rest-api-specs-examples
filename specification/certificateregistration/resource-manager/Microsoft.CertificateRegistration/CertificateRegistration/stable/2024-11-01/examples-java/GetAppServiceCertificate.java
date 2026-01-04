
/**
 * Samples for AppServiceCertificateOrders GetCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/
     * stable/2024-11-01/examples/GetAppServiceCertificate.json
     */
    /**
     * Sample code: Get App Service Certificate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppServiceCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().getCertificateWithResponse(
            "testrg123", "SampleCertificateOrderName", "SampleCertName1", com.azure.core.util.Context.NONE);
    }
}
