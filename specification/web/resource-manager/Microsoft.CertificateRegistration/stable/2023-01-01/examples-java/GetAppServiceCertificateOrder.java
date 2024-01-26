
/**
 * Samples for AppServiceCertificateOrders GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-01-01/examples/
     * GetAppServiceCertificateOrder.json
     */
    /**
     * Sample code: Get App Service Certificate Order.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppServiceCertificateOrder(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().getByResourceGroupWithResponse(
            "testrg123", "SampleCertificateOrderName", com.azure.core.util.Context.NONE);
    }
}
