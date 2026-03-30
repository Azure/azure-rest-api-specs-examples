
/**
 * Samples for Diagnostics GetSiteDetectorResponseSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteDetectorResponseSlot_Slot.json
     */
    /**
     * Sample code: Get App Slot Detector Response.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppSlotDetectorResponse(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteDetectorResponseSlotWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "runtimeavailability", "staging", null, null, null, com.azure.core.util.Context.NONE);
    }
}
