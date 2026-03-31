
/**
 * Samples for Diagnostics GetSiteAnalysis.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteAnalysis.json
     */
    /**
     * Sample code: Get App Analysis.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppAnalysis(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteAnalysisWithResponse("Sample-WestUSResourceGroup", "SampleApp",
            "availability", "appanalysis", com.azure.core.util.Context.NONE);
    }
}
