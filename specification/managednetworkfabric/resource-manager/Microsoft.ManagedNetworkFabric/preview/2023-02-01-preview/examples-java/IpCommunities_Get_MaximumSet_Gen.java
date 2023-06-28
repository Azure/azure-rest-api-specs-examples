/** Samples for IpCommunities GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpCommunities_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpCommunities_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipCommunitiesGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .ipCommunities()
            .getByResourceGroupWithResponse(
                "resourceGroupName", "example-ipCommunity", com.azure.core.util.Context.NONE);
    }
}
