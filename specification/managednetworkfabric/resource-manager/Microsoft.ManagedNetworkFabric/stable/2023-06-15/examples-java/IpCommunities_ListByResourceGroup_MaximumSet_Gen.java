/** Samples for IpCommunities ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpCommunities_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpCommunities_ListByResourceGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipCommunitiesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipCommunities().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
