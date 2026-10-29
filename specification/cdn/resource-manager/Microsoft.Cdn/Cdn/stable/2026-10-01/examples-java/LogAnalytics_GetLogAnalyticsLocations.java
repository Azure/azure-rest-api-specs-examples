
/**
 * Samples for LogAnalytics GetLogAnalyticsLocations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/LogAnalytics_GetLogAnalyticsLocations.json
     */
    /**
     * Sample code: LogAnalytics_GetLogAnalyticsLocations.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void logAnalyticsGetLogAnalyticsLocations(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getLogAnalytics().getLogAnalyticsLocationsWithResponse("RG", "profile1",
            com.azure.core.util.Context.NONE);
    }
}
