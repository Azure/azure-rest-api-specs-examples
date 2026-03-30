
/**
 * Samples for SiteCertificates DeleteSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteSiteCertificateSlot.json
     */
    /**
     * Sample code: Delete Certificate for slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteCertificateForSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getSiteCertificates().deleteSlotWithResponse("testrg123", "testSiteName", "staging",
            "testc6282", com.azure.core.util.Context.NONE);
    }
}
