
/**
 * Samples for WebApps GetNetworkTracesV2.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetWebSiteNetworkTraces_SlotV2.json
     */
    /**
     * Sample code: Get NetworkTraces for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getNetworkTracesForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getNetworkTracesV2WithResponse("testrg123", "SampleApp",
            "c291433b-53ad-4c49-8cae-0a293eae1c6d", com.azure.core.util.Context.NONE);
    }
}
