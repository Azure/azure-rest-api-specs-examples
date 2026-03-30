
/**
 * Samples for Certificates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteCertificate.json
     */
    /**
     * Sample code: Delete Certificate.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteCertificate(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getCertificates().deleteWithResponse("testrg123", "testc6282",
            com.azure.core.util.Context.NONE);
    }
}
