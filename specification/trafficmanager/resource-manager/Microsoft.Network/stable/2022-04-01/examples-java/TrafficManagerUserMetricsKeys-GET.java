/** Samples for TrafficManagerUserMetricsKeys Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/TrafficManagerUserMetricsKeys-GET.json
     */
    /**
     * Sample code: TrafficManagerUserMetricsKeys-GET.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void trafficManagerUserMetricsKeysGET(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getTrafficManagerUserMetricsKeys()
            .getWithResponse(com.azure.core.util.Context.NONE);
    }
}
