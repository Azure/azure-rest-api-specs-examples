
/**
 * Samples for SecurityRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkSecurityGroupRuleDelete.json
     */
    /**
     * Sample code: Delete network security rule from network security group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteNetworkSecurityRuleFromNetworkSecurityGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityRules().delete("rg1", "testnsg", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
