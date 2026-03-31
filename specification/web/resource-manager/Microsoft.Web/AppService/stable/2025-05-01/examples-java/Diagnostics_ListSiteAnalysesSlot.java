
/**
 * Samples for Diagnostics ListSiteAnalyses.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteAnalysesSlot.json
     */
    /**
     * Sample code: List App Slot Analyses.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAppSlotAnalyses(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteAnalyses("Sample-WestUSResourceGroup", "SampleApp",
            "availability", com.azure.core.util.Context.NONE);
    }
}
