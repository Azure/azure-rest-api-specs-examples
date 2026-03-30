
/**
 * Samples for Diagnostics ExecuteSiteDetector.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ExecuteSiteDetectorSlot.json
     */
    /**
     * Sample code: Execute site slot detector.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void executeSiteSlotDetector(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().executeSiteDetectorWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "sitecrashes", "availability", null, null, null, com.azure.core.util.Context.NONE);
    }
}
