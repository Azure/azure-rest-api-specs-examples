
/**
 * Samples for Diagnostics GetSiteDetectorSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteDetector_Slot.json
     */
    /**
     * Sample code: Get App Detector.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppDetector(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteDetectorSlotWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "sitecrashes", "Production", com.azure.core.util.Context.NONE);
    }
}
