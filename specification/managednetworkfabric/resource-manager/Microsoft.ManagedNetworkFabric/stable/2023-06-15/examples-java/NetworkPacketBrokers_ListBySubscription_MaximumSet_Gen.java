/** Samples for NetworkPacketBrokers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkPacketBrokers_ListBySubscription_MaximumSet_Gen.json
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
