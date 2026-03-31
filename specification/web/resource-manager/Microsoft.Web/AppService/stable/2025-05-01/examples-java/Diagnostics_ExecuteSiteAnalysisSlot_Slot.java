
/**
 * Samples for Diagnostics ExecuteSiteAnalysisSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ExecuteSiteAnalysisSlot_Slot.json
     */
    /**
     * Sample code: Execute site slot analysis.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void executeSiteSlotAnalysis(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().executeSiteAnalysisSlotWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "apprestartanalyses", "staging", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
