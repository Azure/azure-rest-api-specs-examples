
/**
 * Samples for Diagnostics GetSiteDetectorResponseSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Diagnostics_GetSiteDetectorResponseSlot.json
     */
    /**
     * Sample code: Get App Slot Detector Response.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppSlotDetectorResponse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().getSiteDetectorResponseSlotWithResponse(
            "Sample-WestUSResourceGroup", "SampleApp", "runtimeavailability", "staging", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
