
/**
 * Samples for AppServiceCertificateOrders List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/
     * ListAppServiceCertificateOrdersBySubscription.json
     */
    /**
     * Sample code: List App Service Certificate orders by subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAppServiceCertificateOrdersBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders()
            .list(com.azure.core.util.Context.NONE);
    }
}
