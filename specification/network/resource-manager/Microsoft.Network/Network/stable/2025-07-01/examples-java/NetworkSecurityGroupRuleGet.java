
/**
 * Samples for SecurityRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityGroupRuleGet.json
     */
    /**
     * Sample code: Get network security rule in network security group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getNetworkSecurityRuleInNetworkSecurityGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityRules().getWithResponse("rg1", "testnsg", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
