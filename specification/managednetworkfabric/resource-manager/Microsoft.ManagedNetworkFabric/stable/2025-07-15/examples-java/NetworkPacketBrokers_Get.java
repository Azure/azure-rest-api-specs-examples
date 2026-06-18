
/**
 * Samples for NetworkPacketBrokers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkPacketBrokers_Get.json
     */
    /**
     * Sample code: NetworkPacketBrokers_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkPacketBrokersGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkPacketBrokers().getByResourceGroupWithResponse("example-rg", "example-networkPacketBroker",
            com.azure.core.util.Context.NONE);
    }
}
