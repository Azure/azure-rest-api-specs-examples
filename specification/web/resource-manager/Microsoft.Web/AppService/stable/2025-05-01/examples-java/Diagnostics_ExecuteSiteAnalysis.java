
/**
 * Samples for Diagnostics ExecuteSiteAnalysis.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ExecuteSiteAnalysis.json
     */
    /**
     * Sample code: Execute site analysis.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void executeSiteAnalysis(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().executeSiteAnalysisWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "apprestartanalyses", null, null, null, com.azure.core.util.Context.NONE);
    }
}
