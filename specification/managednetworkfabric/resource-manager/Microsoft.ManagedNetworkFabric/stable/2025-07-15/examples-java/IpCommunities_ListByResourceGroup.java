
/**
 * Samples for IpCommunities ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpCommunities_ListByResourceGroup.json
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
