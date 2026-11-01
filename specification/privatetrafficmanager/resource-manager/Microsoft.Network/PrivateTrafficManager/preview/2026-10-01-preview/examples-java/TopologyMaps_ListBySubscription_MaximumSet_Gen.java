
/**
 * Samples for TopologyMaps List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/TopologyMaps_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: TopologyMaps_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void topologyMapsListBySubscriptionMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.topologyMaps().list(com.azure.core.util.Context.NONE);
    }
}
