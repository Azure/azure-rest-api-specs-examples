
/**
 * Samples for AppServiceCertificateOrders ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/
     * stable/2024-11-01/examples/ListAppServiceCertificateOrdersByResourceGroup.json
     */
    /**
     * Sample code: List App Service Certificate orders by resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAppServiceCertificateOrdersByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().listByResourceGroup("testrg123",
            com.azure.core.util.Context.NONE);
    }
}
