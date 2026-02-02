
/**
 * Samples for Diagnostics GetSiteAnalysisSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Diagnostics_GetSiteAnalysis.json
     */
    /**
     * Sample code: Get App Analysis.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppAnalysis(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().getSiteAnalysisSlotWithResponse(
            "Sample-WestUSResourceGroup", "SampleApp", "availability", "appanalysis", "Production",
            com.azure.core.util.Context.NONE);
    }
}
