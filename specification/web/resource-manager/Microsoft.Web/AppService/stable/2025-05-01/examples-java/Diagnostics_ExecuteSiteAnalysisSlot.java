
/**
 * Samples for Diagnostics ExecuteSiteAnalysis.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ExecuteSiteAnalysisSlot.json
     */
    /**
     * Sample code: Execute site slot analysis.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void executeSiteSlotAnalysis(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().executeSiteAnalysisWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "apprestartanalyses", null, null, null, com.azure.core.util.Context.NONE);
    }
}
