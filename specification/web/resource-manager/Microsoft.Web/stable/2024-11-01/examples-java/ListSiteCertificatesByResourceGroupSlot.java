
/**
 * Samples for SiteCertificates ListSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/
     * ListSiteCertificatesByResourceGroupSlot.json
     */
    /**
     * Sample code: List Certificates by resource group for a slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCertificatesByResourceGroupForASlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getSiteCertificates().listSlot("testrg123", "testSiteName", "staging",
            com.azure.core.util.Context.NONE);
    }
}
