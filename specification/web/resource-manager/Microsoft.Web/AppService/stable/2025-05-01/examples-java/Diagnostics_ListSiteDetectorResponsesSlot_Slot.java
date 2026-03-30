
/**
 * Samples for Diagnostics ListSiteDetectorResponsesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteDetectorResponsesSlot_Slot.json
     */
    /**
     * Sample code: Get App Slot Detector Responses.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppSlotDetectorResponses(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteDetectorResponsesSlot("Sample-WestUSResourceGroup",
            "SampleApp", "staging", com.azure.core.util.Context.NONE);
    }
}
