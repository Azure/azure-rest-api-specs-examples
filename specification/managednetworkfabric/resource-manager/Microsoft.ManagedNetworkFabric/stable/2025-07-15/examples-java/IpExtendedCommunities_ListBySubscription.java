
/**
 * Samples for IpExtendedCommunities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpExtendedCommunities_ListBySubscription.json
     */
    /**
     * Sample code: IpExtendedCommunities_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipExtendedCommunitiesListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipExtendedCommunities().list(com.azure.core.util.Context.NONE);
    }
}
