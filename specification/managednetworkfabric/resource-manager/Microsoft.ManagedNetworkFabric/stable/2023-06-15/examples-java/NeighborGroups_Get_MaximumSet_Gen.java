/** Samples for NeighborGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NeighborGroups_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NeighborGroups_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void neighborGroupsGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .neighborGroups()
            .getByResourceGroupWithResponse("example-rg", "example-neighborGroup", com.azure.core.util.Context.NONE);
    }
}
