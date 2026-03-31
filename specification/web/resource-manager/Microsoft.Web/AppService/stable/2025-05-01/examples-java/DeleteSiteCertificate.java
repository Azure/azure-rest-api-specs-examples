
/**
 * Samples for SiteCertificates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteSiteCertificate.json
     */
    /**
     * Sample code: Delete Certificate.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteCertificate(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getSiteCertificates().deleteWithResponse("testrg123", "testSiteName", "testc6282",
            com.azure.core.util.Context.NONE);
    }
}
