
/**
 * Samples for LogAnalytics GetLogAnalyticsResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/LogAnalytics_GetLogAnalyticsResources.json
     */
    /**
     * Sample code: LogAnalytics_GetLogAnalyticsResources.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void logAnalyticsGetLogAnalyticsResources(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getLogAnalytics().getLogAnalyticsResourcesWithResponse("RG", "profile1",
            com.azure.core.util.Context.NONE);
    }
}
