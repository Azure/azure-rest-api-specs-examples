
/**
 * Samples for Diagnostics ExecuteSiteAnalysisSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ExecuteSiteAnalysis_Slot.json
     */
    /**
     * Sample code: Execute site analysis.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void executeSiteAnalysis(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().executeSiteAnalysisSlotWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "apprestartanalyses", "Production", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
