
/**
 * Samples for Diagnostics ListSiteAnalysesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/Diagnostics_ListSiteAnalyses.json
     */
    /**
     * Sample code: List App Analyses.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppAnalyses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().listSiteAnalysesSlot("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "Production", com.azure.core.util.Context.NONE);
    }
}
