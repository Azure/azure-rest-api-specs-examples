
/**
 * Samples for AppServiceCertificateOrders RetrieveCertificateActions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/
     * RetrieveCertificateOrderActions.json
     */
    /**
     * Sample code: Retrieve Certificate Order Actions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveCertificateOrderActions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders()
            .retrieveCertificateActionsWithResponse("testrg123", "SampleCertOrder", com.azure.core.util.Context.NONE);
    }
}
