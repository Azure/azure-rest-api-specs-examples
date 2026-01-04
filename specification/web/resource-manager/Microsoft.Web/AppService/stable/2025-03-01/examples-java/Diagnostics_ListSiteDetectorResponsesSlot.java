
/**
 * Samples for Diagnostics ListSiteDetectorResponses.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * Diagnostics_ListSiteDetectorResponsesSlot.json
     */
    /**
     * Sample code: Get App Slot Detector Responses.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppSlotDetectorResponses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDiagnostics()
            .listSiteDetectorResponses("Sample-WestUSResourceGroup", "SampleApp", com.azure.core.util.Context.NONE);
    }
}
