
/**
 * Samples for NetworkSecurityGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityGroupDelete.json
     */
    /**
     * Sample code: Delete network security group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteNetworkSecurityGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityGroups().delete("rg1", "testnsg", com.azure.core.util.Context.NONE);
    }
}
