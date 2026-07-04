
/**
 * Samples for SecurityRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityGroupRuleDelete.json
     */
    /**
     * Sample code: Delete network security rule from network security group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteNetworkSecurityRuleFromNetworkSecurityGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityRules().delete("rg1", "testnsg", "rule1", com.azure.core.util.Context.NONE);
    }
}
