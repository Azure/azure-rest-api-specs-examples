
/**
 * Samples for Diagnostics GetSiteAnalysis.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteAnalysisSlot.json
     */
    /**
     * Sample code: Get App Slot Analysis.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppSlotAnalysis(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteAnalysisWithResponse("Sample-WestUSResourceGroup", "SampleApp",
            "availability", "appanalysis", com.azure.core.util.Context.NONE);
    }
}
