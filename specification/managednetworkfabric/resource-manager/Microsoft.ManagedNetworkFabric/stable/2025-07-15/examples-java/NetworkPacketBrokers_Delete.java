
/**
 * Samples for NetworkPacketBrokers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkPacketBrokers_Delete.json
     */
    /**
     * Sample code: NetworkPacketBrokers_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkPacketBrokersDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkPacketBrokers().delete("example-rg", "example-networkPacketBroker",
            com.azure.core.util.Context.NONE);
    }
}
