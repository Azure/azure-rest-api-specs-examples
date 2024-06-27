
import com.azure.resourcemanager.appservice.models.ReissueCertificateOrderRequest;

/**
 * Samples for AppServiceCertificateOrders Reissue.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/
     * ReissueAppServiceCertificateOrder.json
     */
    /**
     * Sample code: Reissue App Service Certificate Order.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reissueAppServiceCertificateOrder(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().reissueWithResponse("testrg123",
            "SampleCertificateOrderName", new ReissueCertificateOrderRequest().withKeySize(2048)
                .withDelayExistingRevokeInHours(2).withCsr("CSR1223238Value").withIsPrivateKeyExternal(false),
            com.azure.core.util.Context.NONE);
    }
}
