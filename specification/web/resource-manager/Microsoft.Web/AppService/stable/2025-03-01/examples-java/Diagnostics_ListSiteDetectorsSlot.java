
/**
 * Samples for Diagnostics ListSiteDetectors.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Diagnostics_ListSiteDetectorsSlot.json
     */
    /**
     * Sample code: List App Slot Detectors.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppSlotDetectors(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().listSiteDetectors("Sample-WestUSResourceGroup",
            "SampleApp", "availability", com.azure.core.util.Context.NONE);
    }
}
