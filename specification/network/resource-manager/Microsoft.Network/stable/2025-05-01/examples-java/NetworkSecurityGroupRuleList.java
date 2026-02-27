
/**
 * Samples for SecurityRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkSecurityGroupRuleList.
     * json
     */
    /**
     * Sample code: List network security rules in network security group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listNetworkSecurityRulesInNetworkSecurityGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityRules().list("rg1", "testnsg",
            com.azure.core.util.Context.NONE);
    }
}
