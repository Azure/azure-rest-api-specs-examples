
/**
 * Samples for Diagnostics ListSiteAnalyses.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/Diagnostics_ListSiteAnalysesSlot.json
     */
    /**
     * Sample code: List App Slot Analyses.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppSlotAnalyses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().listSiteAnalyses("Sample-WestUSResourceGroup",
            "SampleApp", "availability", com.azure.core.util.Context.NONE);
    }
}
