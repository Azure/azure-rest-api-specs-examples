import com.azure.core.util.Context;

/** Samples for Diagnostics GetSiteAnalysis. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_GetSiteAnalysisSlot.json
     */
    /**
     * Sample code: Get App Slot Analysis.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppSlotAnalysis(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .getSiteAnalysisWithResponse(
                "Sample-WestUSResourceGroup", "SampleApp", "availability", "appanalysis", Context.NONE);
    }
}
