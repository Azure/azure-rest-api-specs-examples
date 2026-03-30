
/**
 * Samples for WebApps StopNetworkTraceSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/StopWebSiteNetworkTrace_StopNetworkTraceSlot.json
     */
    /**
     * Sample code: Stop a currently running network trace operation for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void stopACurrentlyRunningNetworkTraceOperationForASite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().stopNetworkTraceSlotWithResponse("testrg123", "SampleApp", "Production",
            com.azure.core.util.Context.NONE);
    }
}
