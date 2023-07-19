/** Samples for NetworkPacketBrokers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkPacketBrokers_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkPacketBrokers_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkPacketBrokersGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkPacketBrokers()
            .getByResourceGroupWithResponse(
                "example-rg", "example-networkPacketBroker", com.azure.core.util.Context.NONE);
    }
}
