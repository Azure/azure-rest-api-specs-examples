/** Samples for NeighborGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NeighborGroups_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NeighborGroups_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void neighborGroupsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.neighborGroups().delete("example-rg", "example-neighborGroup", com.azure.core.util.Context.NONE);
    }
}
