
/**
 * Samples for Diagnostics GetSiteDetectorResponse.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteDetectorResponseSlot.json
     */
    /**
     * Sample code: Get App Slot Detector Response.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppSlotDetectorResponse(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteDetectorResponseWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "runtimeavailability", null, null, null, com.azure.core.util.Context.NONE);
    }
}
