
/**
 * Samples for WebApps StartWebSiteNetworkTraceOperationSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/StartWebSiteNetworkTraceOperation.
     * json
     */
    /**
     * Sample code: Start a new network trace operation for a site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startANewNetworkTraceOperationForASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().startWebSiteNetworkTraceOperationSlot("testrg123",
            "SampleApp", "Production", 60, null, null, com.azure.core.util.Context.NONE);
    }
}
