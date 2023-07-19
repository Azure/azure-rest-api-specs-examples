/** Samples for IpCommunities Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpCommunities_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpCommunities_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipCommunitiesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipCommunities().delete("example-rg", "example-ipcommunity", com.azure.core.util.Context.NONE);
    }
}
