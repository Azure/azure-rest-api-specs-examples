/** Samples for TrafficManagerUserMetricsKeys Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/TrafficManagerUserMetricsKeys-DELETE.json
     */
    /**
     * Sample code: TrafficManagerUserMetricsKeys-DELETE.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void trafficManagerUserMetricsKeysDELETE(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getTrafficManagerUserMetricsKeys()
            .deleteWithResponse(com.azure.core.util.Context.NONE);
    }
}
