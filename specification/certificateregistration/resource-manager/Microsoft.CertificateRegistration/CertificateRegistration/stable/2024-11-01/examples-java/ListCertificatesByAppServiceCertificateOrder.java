
/**
 * Samples for AppServiceCertificateOrders ListCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/
     * stable/2024-11-01/examples/ListCertificatesByAppServiceCertificateOrder.json
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
