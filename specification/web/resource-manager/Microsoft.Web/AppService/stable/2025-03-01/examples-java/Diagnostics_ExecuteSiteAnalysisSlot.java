
/**
 * Samples for Diagnostics ExecuteSiteAnalysisSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Diagnostics_ExecuteSiteAnalysisSlot.json
     */
    /**
     * Sample code: Execute site slot analysis.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void executeSiteSlotAnalysis(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().executeSiteAnalysisSlotWithResponse(
            "Sample-WestUSResourceGroup", "SampleApp", "availability", "apprestartanalyses", "staging", null, null,
            null, com.azure.core.util.Context.NONE);
    }
}
