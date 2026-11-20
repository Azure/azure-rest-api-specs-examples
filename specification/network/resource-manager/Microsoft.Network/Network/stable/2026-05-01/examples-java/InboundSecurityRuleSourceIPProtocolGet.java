
/**
 * Samples for InboundSecurityRuleOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/InboundSecurityRuleSourceIPProtocolGet.json
     */
    /**
     * Sample code: Get a Permanent Internet Ingress rule with SourceIPProtocol load distribution.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAPermanentInternetIngressRuleWithSourceIPProtocolLoadDistribution(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getInboundSecurityRuleOperations().getWithResponse("rg1", "nva", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
