
import com.azure.resourcemanager.appservice.models.RenewCertificateOrderRequest;

/**
 * Samples for AppServiceCertificateOrders Renew.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/
     * stable/2024-11-01/examples/RenewAppServiceCertificateOrder.json
     */
    /**
     * Sample code: Renew App Service Certificate Order.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void renewAppServiceCertificateOrder(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders()
            .renewWithResponse("testrg123", "SampleCertificateOrderName", new RenewCertificateOrderRequest()
                .withKeySize(2048).withCsr("CSR1223238Value").withIsPrivateKeyExternal(false),
                com.azure.core.util.Context.NONE);
    }
}
