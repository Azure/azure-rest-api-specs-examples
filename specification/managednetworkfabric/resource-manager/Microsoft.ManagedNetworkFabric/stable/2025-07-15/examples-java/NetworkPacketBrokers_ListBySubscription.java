
/**
 * Samples for NetworkPacketBrokers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkPacketBrokers_ListBySubscription.json
     */
    /**
     * Sample code: NetworkPacketBrokers_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkPacketBrokersListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkPacketBrokers().list(com.azure.core.util.Context.NONE);
    }
}
