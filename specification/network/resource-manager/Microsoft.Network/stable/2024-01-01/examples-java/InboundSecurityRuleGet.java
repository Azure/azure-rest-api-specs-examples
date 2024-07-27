
/**
 * Samples for InboundSecurityRuleOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/InboundSecurityRuleGet.json
     */
    /**
     * Sample code: Create Network Virtual Appliance Inbound Security Rules.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createNetworkVirtualApplianceInboundSecurityRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getInboundSecurityRuleOperations().getWithResponse("rg1", "nva",
            "rule1", com.azure.core.util.Context.NONE);
    }
}
