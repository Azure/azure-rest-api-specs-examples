
/**
 * Samples for SiteCertificates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/DeleteSiteCertificate.json
     */
    /**
     * Sample code: Delete Certificate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getSiteCertificates().deleteWithResponse("testrg123", "testSiteName",
            "testc6282", com.azure.core.util.Context.NONE);
    }
}
