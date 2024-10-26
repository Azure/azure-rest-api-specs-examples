
/**
 * Samples for Diagnostics ExecuteSiteDetectorSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/Diagnostics_ExecuteSiteDetector.json
     */
    /**
     * Sample code: Execute site detector.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void executeSiteDetector(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().executeSiteDetectorSlotWithResponse(
            "Sample-WestUSResourceGroup", "SampleApp", "sitecrashes", "availability", "Production", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
