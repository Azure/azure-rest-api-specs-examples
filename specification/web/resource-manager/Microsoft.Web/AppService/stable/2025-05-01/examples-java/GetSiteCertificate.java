
/**
 * Samples for SiteCertificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetSiteCertificate.json
     */
    /**
     * Sample code: Get Certificate.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getCertificate(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getSiteCertificates().getWithResponse("testrg123", "testSiteName", "testc6282",
            com.azure.core.util.Context.NONE);
    }
}
