/** Samples for IpExtendedCommunities GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpExtendedCommunities_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipExtendedCommunitiesGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .ipExtendedCommunities()
            .getByResourceGroupWithResponse(
                "example-rg", "example-ipExtendedCommunity", com.azure.core.util.Context.NONE);
    }
}
