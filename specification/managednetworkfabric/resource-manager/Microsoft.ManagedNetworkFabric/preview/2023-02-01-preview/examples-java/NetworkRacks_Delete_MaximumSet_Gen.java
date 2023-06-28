/** Samples for NetworkRacks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkRacks_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkRacks_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkRacksDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkRacks().delete("resourceGroupName", "networkRackName", com.azure.core.util.Context.NONE);
    }
}
