
/**
 * Samples for WebApps StartNetworkTrace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/StartWebSiteNetworkTraceOperation_StartNetworkTrace.json
     */
    /**
     * Sample code: Start a new network trace operation for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        startANewNetworkTraceOperationForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().startNetworkTrace("testrg123", "SampleApp", 60, null, null,
            com.azure.core.util.Context.NONE);
    }
}
