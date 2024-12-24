
/**
 * Samples for Diagnostics ExecuteSiteAnalysis.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/Diagnostics_ExecuteSiteAnalysisSlot.
     * json
     */
    /**
     * Sample code: Execute site slot analysis.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void executeSiteSlotAnalysis(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().executeSiteAnalysisWithResponse(
            "Sample-WestUSResourceGroup", "SampleApp", "availability", "apprestartanalyses", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
