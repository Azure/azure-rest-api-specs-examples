/** Samples for IpExtendedCommunities Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpExtendedCommunities_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpExtendedCommunities_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipExtendedCommunitiesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .ipExtendedCommunities()
            .delete("rgIpExtendedCommunityLists", "example_ipExtendedCommunity", com.azure.core.util.Context.NONE);
    }
}
