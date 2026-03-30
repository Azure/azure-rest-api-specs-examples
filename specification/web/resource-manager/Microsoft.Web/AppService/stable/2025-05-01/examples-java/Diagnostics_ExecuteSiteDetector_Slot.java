
/**
 * Samples for Diagnostics ExecuteSiteDetectorSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ExecuteSiteDetector_Slot.json
     */
    /**
     * Sample code: Execute site detector.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void executeSiteDetector(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().executeSiteDetectorSlotWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "sitecrashes", "availability", "Production", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
