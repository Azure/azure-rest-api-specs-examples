
/**
 * Samples for Diagnostics GetSiteAnalysisSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/Diagnostics_GetSiteAnalysis_Slot.json
     */
    /**
     * Sample code: Get App Analysis.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppAnalysis(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getDiagnostics().getSiteAnalysisSlotWithResponse("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "appanalysis", "Production", com.azure.core.util.Context.NONE);
    }
}
