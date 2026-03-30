
/**
 * Samples for Diagnostics GetSiteDetectorSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteDetectorSlot_Slot.json
     */
    /**
     * Sample code: Get App Slot Detector.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppSlotDetector(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteDetectorSlotWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "sitecrashes", "staging", com.azure.core.util.Context.NONE);
    }
}
