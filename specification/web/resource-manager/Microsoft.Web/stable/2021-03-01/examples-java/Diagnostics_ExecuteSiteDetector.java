import com.azure.core.util.Context;

/** Samples for Diagnostics ExecuteSiteDetector. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_ExecuteSiteDetector.json
     */
    /**
     * Sample code: Execute site detector.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void executeSiteDetector(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .executeSiteDetectorWithResponse(
                "Sample-WestUSResourceGroup",
                "SampleApp",
                "sitecrashes",
                "availability",
                null,
                null,
                null,
                Context.NONE);
    }
}
