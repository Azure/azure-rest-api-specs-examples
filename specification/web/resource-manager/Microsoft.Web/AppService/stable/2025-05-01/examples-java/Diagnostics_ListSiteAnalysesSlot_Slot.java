
/**
 * Samples for Diagnostics ListSiteAnalysesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteAnalysesSlot_Slot.json
     */
    /**
     * Sample code: List App Slot Analyses.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAppSlotAnalyses(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().listSiteAnalysesSlot("Sample-WestUSResourceGroup", "SampleApp",
            "availability", "staging", com.azure.core.util.Context.NONE);
    }
}
