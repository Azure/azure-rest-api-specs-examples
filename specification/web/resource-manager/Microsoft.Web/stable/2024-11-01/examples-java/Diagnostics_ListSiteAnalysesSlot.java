
/**
 * Samples for Diagnostics ListSiteAnalysesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/Diagnostics_ListSiteAnalysesSlot.json
     */
    /**
     * Sample code: List App Slot Analyses.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppSlotAnalyses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().listSiteAnalysesSlot("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "staging", com.azure.core.util.Context.NONE);
    }
}
