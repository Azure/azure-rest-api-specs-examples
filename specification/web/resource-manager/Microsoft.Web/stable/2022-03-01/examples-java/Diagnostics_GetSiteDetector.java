import com.azure.core.util.Context;

/** Samples for Diagnostics GetSiteDetector. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/Diagnostics_GetSiteDetector.json
     */
    /**
     * Sample code: Get App Detector.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppDetector(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDiagnostics()
            .getSiteDetectorWithResponse(
                "Sample-WestUSResourceGroup", "SampleApp", "availability", "sitecrashes", Context.NONE);
    }
}
