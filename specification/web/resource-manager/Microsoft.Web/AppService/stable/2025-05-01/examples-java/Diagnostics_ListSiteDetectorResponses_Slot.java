
/**
 * Samples for Diagnostics ListSiteDetectorResponsesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteDetectorResponses_Slot.json
     */
    /**
     * Sample code: Get App Detector Responses.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppDetectorResponses(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteDetectorResponsesSlot("Sample-WestUSResourceGroup",
            "SampleApp", "staging", com.azure.core.util.Context.NONE);
    }
}
