/** Samples for IpCommunities Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpCommunities_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: IpCommunities_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipCommunitiesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipCommunities().delete("resourceGroupName", "IpCommunityList1", com.azure.core.util.Context.NONE);
    }
}
