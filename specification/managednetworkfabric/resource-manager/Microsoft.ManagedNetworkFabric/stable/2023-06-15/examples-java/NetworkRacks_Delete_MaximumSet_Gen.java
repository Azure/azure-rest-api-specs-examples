/** Samples for NetworkRacks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkRacks_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkRacks_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkRacksDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkRacks().delete("example-rg", "example-rack", com.azure.core.util.Context.NONE);
    }
}
