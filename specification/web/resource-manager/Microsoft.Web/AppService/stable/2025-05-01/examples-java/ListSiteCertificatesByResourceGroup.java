
/**
 * Samples for SiteCertificates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListSiteCertificatesByResourceGroup.json
     */
    /**
     * Sample code: List Certificates by resource group.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listCertificatesByResourceGroup(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getSiteCertificates().list("testrg123", "testSiteName",
            com.azure.core.util.Context.NONE);
    }
}
