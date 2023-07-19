/** Samples for IpExtendedCommunities Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_Delete_MaximumSet_Gen.json
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
            .delete("example-rg", "example-ipExtendedCommunity", com.azure.core.util.Context.NONE);
    }
}
