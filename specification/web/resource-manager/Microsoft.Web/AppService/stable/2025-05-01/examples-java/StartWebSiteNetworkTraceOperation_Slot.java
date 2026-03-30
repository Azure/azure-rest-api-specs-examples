
/**
 * Samples for WebApps StartWebSiteNetworkTraceOperationSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/StartWebSiteNetworkTraceOperation_Slot.json
     */
    /**
     * Sample code: Start a new network trace operation for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        startANewNetworkTraceOperationForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().startWebSiteNetworkTraceOperationSlot("testrg123", "SampleApp",
            "Production", 60, null, null, com.azure.core.util.Context.NONE);
    }
}
