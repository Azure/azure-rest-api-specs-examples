/** Samples for IpExtendedCommunities ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpExtendedCommunities_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpExtendedCommunities_ListByResourceGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipExtendedCommunitiesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .ipExtendedCommunities()
            .listByResourceGroup("rgIpExtendedCommunityLists", com.azure.core.util.Context.NONE);
    }
}
