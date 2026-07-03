
/**
 * Samples for NetworkSecurityGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityGroupGet.json
     */
    /**
     * Sample code: Get network security group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getNetworkSecurityGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityGroups().getByResourceGroupWithResponse("rg1", "testnsg", null,
            com.azure.core.util.Context.NONE);
    }
}
