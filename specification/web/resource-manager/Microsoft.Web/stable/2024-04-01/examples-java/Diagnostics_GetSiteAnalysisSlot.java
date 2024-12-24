
/**
 * Samples for Diagnostics GetSiteAnalysisSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/Diagnostics_GetSiteAnalysisSlot.json
     */
    /**
     * Sample code: Get App Slot Analysis.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppSlotAnalysis(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().getSiteAnalysisSlotWithResponse(
            "Sample-WestUSResourceGroup", "SampleApp", "availability", "appanalysis", "staging",
            com.azure.core.util.Context.NONE);
    }
}
