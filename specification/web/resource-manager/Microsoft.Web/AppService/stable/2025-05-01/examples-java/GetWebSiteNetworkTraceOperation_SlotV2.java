
/**
 * Samples for WebApps GetNetworkTraceOperationSlotV2.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetWebSiteNetworkTraceOperation_SlotV2.json
     */
    /**
     * Sample code: Get the current status of a network trace operation for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getTheCurrentStatusOfANetworkTraceOperationForASite(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getNetworkTraceOperationSlotV2WithResponse("testrg123", "SampleApp",
            "c291433b-53ad-4c49-8cae-0a293eae1c6d", "Production", com.azure.core.util.Context.NONE);
    }
}
