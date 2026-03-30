
/**
 * Samples for Diagnostics ListSiteAnalyses.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteAnalyses.json
     */
    /**
     * Sample code: List App Analyses.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAppAnalyses(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteAnalyses("Sample-WestUSResourceGroup", "SampleApp",
            "availability", com.azure.core.util.Context.NONE);
    }
}
