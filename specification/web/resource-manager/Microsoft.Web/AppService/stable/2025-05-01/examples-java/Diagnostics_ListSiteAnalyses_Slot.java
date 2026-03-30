
/**
 * Samples for Diagnostics ListSiteAnalysesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteAnalyses_Slot.json
     */
    /**
     * Sample code: List App Analyses.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAppAnalyses(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteAnalysesSlot("Sample-WestUSResourceGroup", "SampleApp",
            "availability", "Production", com.azure.core.util.Context.NONE);
    }
}
