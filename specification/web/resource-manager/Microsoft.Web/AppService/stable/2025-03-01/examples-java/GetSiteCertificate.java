
/**
 * Samples for SiteCertificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetSiteCertificate.json
     */
    /**
     * Sample code: Get Certificate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getSiteCertificates().getWithResponse("testrg123", "testSiteName",
            "testc6282", com.azure.core.util.Context.NONE);
    }
}
