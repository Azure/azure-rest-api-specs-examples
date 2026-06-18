
/**
 * Samples for IpCommunities GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpCommunities_Get.json
     */
    /**
     * Sample code: IpCommunities_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipCommunitiesGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipCommunities().getByResourceGroupWithResponse("example-rg", "example-ipcommunity",
            com.azure.core.util.Context.NONE);
    }
}
