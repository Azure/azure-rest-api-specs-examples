
/**
 * Samples for Certificates ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ListCertificatesByResourceGroup.json
     */
    /**
     * Sample code: List Certificates by resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCertificatesByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getCertificates().listByResourceGroup("testrg123",
            com.azure.core.util.Context.NONE);
    }
}
