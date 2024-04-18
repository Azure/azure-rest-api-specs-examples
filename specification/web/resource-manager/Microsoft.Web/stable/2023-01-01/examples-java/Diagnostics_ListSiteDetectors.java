
/**
 * Samples for Diagnostics ListSiteDetectors.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/Diagnostics_ListSiteDetectors.json
     */
    /**
     * Sample code: List App Detectors.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppDetectors(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().listSiteDetectors("Sample-WestUSResourceGroup",
            "SampleApp", "availability", com.azure.core.util.Context.NONE);
    }
}
