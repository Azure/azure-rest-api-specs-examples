import com.azure.core.util.Context;

/** Samples for LogAnalytics GetLogAnalyticsResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsResources.json
     */
    /**
     * Sample code: LogAnalytics_GetLogAnalyticsResources.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void logAnalyticsGetLogAnalyticsResources(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getLogAnalytics()
            .getLogAnalyticsResourcesWithResponse("RG", "profile1", Context.NONE);
    }
}
