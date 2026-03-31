
/**
 * Samples for WebApps StopWebSiteNetworkTrace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/StopWebSiteNetworkTrace.json
     */
    /**
     * Sample code: Stop a currently running network trace operation for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void stopACurrentlyRunningNetworkTraceOperationForASite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().stopWebSiteNetworkTraceWithResponse("testrg123", "SampleApp",
            com.azure.core.util.Context.NONE);
    }
}
