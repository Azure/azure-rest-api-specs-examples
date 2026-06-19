
/**
 * Samples for NetworkPacketBrokers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkPacketBrokers_ListByResourceGroup.json
     */
    /**
     * Sample code: NetworkPacketBrokers_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkPacketBrokersListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkPacketBrokers().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
