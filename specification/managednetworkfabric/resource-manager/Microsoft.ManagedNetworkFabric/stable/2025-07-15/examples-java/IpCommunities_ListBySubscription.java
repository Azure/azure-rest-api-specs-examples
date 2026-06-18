
/**
 * Samples for IpCommunities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpCommunities_ListBySubscription.json
     */
    /**
     * Sample code: IpCommunities_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipCommunitiesListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipCommunities().list(com.azure.core.util.Context.NONE);
    }
}
