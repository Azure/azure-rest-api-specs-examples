/** Samples for TrafficManagerUserMetricsKeys CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/TrafficManagerUserMetricsKeys-PUT.json
     */
    /**
     * Sample code: TrafficManagerUserMetricsKeys-PUT.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void trafficManagerUserMetricsKeysPUT(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getTrafficManagerUserMetricsKeys()
            .createOrUpdateWithResponse(com.azure.core.util.Context.NONE);
    }
}
