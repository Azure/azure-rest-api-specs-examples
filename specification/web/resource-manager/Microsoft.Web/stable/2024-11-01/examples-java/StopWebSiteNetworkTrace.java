
/**
 * Samples for WebApps StopNetworkTrace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/StopWebSiteNetworkTrace.json
     */
    /**
     * Sample code: Stop a currently running network trace operation for a site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        stopACurrentlyRunningNetworkTraceOperationForASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().stopNetworkTraceWithResponse("testrg123", "SampleApp",
            com.azure.core.util.Context.NONE);
    }
}
