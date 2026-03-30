
/**
 * Samples for Diagnostics GetSiteAnalysisSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteAnalysisSlot_Slot.json
     */
    /**
     * Sample code: Get App Slot Analysis.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppSlotAnalysis(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteAnalysisSlotWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "appanalysis", "staging", com.azure.core.util.Context.NONE);
    }
}
