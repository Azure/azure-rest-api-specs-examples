
/**
 * Samples for AppServiceCertificateOrders ListCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/
     * ListCertificatesByAppServiceCertificateOrder.json
     */
    /**
     * Sample code: List certificates by App Service Certificate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCertificatesByAppServiceCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().listCertificates("testrg123",
            "SampleCertificateOrderName", com.azure.core.util.Context.NONE);
    }
}
