
/**
 * Samples for Diagnostics ListSiteDetectorsSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/Diagnostics_ListSiteDetectorsSlot.
     * json
     */
    /**
     * Sample code: List App Slot Detectors.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppSlotDetectors(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics().listSiteDetectorsSlot("Sample-WestUSResourceGroup",
            "SampleApp", "availability", "staging", com.azure.core.util.Context.NONE);
    }
}
