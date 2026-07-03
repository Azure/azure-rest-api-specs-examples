
/**
 * Samples for InboundSecurityRuleOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/InboundSecurityRuleGet.json
     */
    /**
     * Sample code: Create Network Virtual Appliance Inbound Security Rules.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createNetworkVirtualApplianceInboundSecurityRules(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getInboundSecurityRuleOperations().getWithResponse("rg1", "nva", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
