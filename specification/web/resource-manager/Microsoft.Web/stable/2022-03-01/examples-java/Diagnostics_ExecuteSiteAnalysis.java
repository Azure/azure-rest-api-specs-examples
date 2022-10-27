import com.azure.core.util.Context;

/** Samples for Diagnostics ExecuteSiteAnalysis. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/Diagnostics_ExecuteSiteAnalysis.json
     */
    /**
     * Sample code: Execute site analysis.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void executeSiteAnalysis(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .executeSiteAnalysisWithResponse(
                "Sample-WestUSResourceGroup",
                "SampleApp",
                "availability",
                "apprestartanalyses",
                null,
                null,
                null,
                Context.NONE);
    }
}
