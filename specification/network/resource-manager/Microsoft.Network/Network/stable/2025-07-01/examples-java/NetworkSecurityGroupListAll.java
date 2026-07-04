
/**
 * Samples for NetworkSecurityGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityGroupListAll.json
     */
    /**
     * Sample code: List all network security groups.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllNetworkSecurityGroups(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityGroups().list(com.azure.core.util.Context.NONE);
    }
}
