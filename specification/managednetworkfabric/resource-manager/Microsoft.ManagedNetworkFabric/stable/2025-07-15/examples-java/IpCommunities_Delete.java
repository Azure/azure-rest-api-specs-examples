
/**
 * Samples for IpCommunities Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/IpCommunities_Delete.json
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
