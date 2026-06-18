
/**
 * Samples for IpExtendedCommunities Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpExtendedCommunities_Delete.json
     */
    /**
     * Sample code: IpExtendedCommunities_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void ipExtendedCommunitiesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.ipExtendedCommunities().delete("example-rg", "example-ipExtendedCommunity",
            com.azure.core.util.Context.NONE);
    }
}
