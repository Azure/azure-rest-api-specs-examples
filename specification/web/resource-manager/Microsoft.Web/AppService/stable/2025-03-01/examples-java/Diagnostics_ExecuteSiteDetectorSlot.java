
/**
 * Samples for Diagnostics ExecuteSiteDetectorSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Diagnostics_ExecuteSiteDetectorSlot.json
     */
    /**
     * Sample code: Execute site slot detector.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void executeSiteSlotDetector(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().executeSiteDetectorSlotWithResponse(
            "Sample-WestUSResourceGroup", "SampleApp", "sitecrashes", "availability", "staging", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
