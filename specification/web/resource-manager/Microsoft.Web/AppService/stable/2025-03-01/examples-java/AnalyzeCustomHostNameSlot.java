
/**
 * Samples for WebApps AnalyzeCustomHostnameSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/AnalyzeCustomHostNameSlot.
     * json
     */
    /**
     * Sample code: Analyze custom hostname for slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void analyzeCustomHostnameForSlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().analyzeCustomHostnameSlotWithResponse("testrg123",
            "sitef6141", "staging", null, com.azure.core.util.Context.NONE);
    }
}
