
/**
 * Samples for AppServiceCertificateOrders RetrieveCertificateEmailHistory.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/
     * RetrieveCertificateEmailHistory.json
     */
    /**
     * Sample code: Retrieve Certificate Email History.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveCertificateEmailHistory(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders()
            .retrieveCertificateEmailHistoryWithResponse("testrg123", "SampleCertOrder",
                com.azure.core.util.Context.NONE);
    }
}
