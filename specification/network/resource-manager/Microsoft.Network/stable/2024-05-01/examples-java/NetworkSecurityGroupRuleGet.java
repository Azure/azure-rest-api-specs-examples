
/**
 * Samples for SecurityRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkSecurityGroupRuleGet.
     * json
     */
    /**
     * Sample code: Get network security rule in network security group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getNetworkSecurityRuleInNetworkSecurityGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityRules().getWithResponse("rg1", "testnsg", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
