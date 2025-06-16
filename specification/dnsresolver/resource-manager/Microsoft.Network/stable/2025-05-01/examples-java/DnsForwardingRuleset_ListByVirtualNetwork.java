
/**
 * Samples for DnsForwardingRulesets ListByVirtualNetwork.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * DnsForwardingRuleset_ListByVirtualNetwork.json
     */
    /**
     * Sample code: List DNS forwarding rulesets by virtual network.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSForwardingRulesetsByVirtualNetwork(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsForwardingRulesets().listByVirtualNetwork("sampleResourceGroup", "sampleVirtualNetwork", null,
            com.azure.core.util.Context.NONE);
    }
}
