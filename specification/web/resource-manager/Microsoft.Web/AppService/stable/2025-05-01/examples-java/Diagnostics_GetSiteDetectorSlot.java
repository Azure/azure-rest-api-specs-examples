
/**
 * Samples for Diagnostics GetSiteDetector.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteDetectorSlot.json
     */
    /**
     * Sample code: Get App Slot Detector.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppSlotDetector(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteDetectorWithResponse("Sample-WestUSResourceGroup", "SampleApp",
            "availability", "sitecrashes", com.azure.core.util.Context.NONE);
    }
}
