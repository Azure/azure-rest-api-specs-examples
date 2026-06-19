
/**
 * Samples for IpExtendedCommunities ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpExtendedCommunities_ListByResourceGroup.json
     */
    /**
     * Sample code: IpExtendedCommunities_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipExtendedCommunitiesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipExtendedCommunities().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
