
/**
 * Samples for Certificates List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListCertificates.json
     */
    /**
     * Sample code: List Certificates for subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCertificatesForSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getCertificates().list(null, com.azure.core.util.Context.NONE);
    }
}
