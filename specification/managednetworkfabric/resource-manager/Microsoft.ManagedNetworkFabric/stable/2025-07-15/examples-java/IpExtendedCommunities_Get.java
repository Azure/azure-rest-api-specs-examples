
/**
 * Samples for IpExtendedCommunities GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpExtendedCommunities_Get.json
     */
    /**
     * Sample code: IpExtendedCommunities_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipExtendedCommunitiesGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipExtendedCommunities().getByResourceGroupWithResponse("example-rg", "example-ipExtendedCommunity",
            com.azure.core.util.Context.NONE);
    }
}
