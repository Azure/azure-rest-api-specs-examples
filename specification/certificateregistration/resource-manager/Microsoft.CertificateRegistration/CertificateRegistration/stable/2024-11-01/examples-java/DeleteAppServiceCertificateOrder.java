
/**
 * Samples for AppServiceCertificateOrders Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/
     * stable/2024-11-01/examples/DeleteAppServiceCertificateOrder.json
     */
    /**
     * Sample code: Delete App Service Certificate Order.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAppServiceCertificateOrder(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().deleteWithResponse("testrg123",
            "SampleCertificateOrderName", com.azure.core.util.Context.NONE);
    }
}
