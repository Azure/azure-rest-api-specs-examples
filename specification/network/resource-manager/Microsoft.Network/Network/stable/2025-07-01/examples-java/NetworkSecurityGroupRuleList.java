
/**
 * Samples for SecurityRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityGroupRuleList.json
     */
    /**
     * Sample code: List network security rules in network security group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listNetworkSecurityRulesInNetworkSecurityGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityRules().list("rg1", "testnsg", com.azure.core.util.Context.NONE);
    }
}
