import com.azure.core.util.Context;

/** Samples for LogAnalytics GetLogAnalyticsLocations. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsLocations.json
     */
    /**
     * Sample code: LogAnalytics_GetLogAnalyticsLocations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void logAnalyticsGetLogAnalyticsLocations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getLogAnalytics()
            .getLogAnalyticsLocationsWithResponse("RG", "profile1", Context.NONE);
    }
}
