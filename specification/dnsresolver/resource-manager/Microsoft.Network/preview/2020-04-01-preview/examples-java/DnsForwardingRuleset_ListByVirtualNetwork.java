import com.azure.core.util.Context;

/** Samples for DnsForwardingRulesets ListByVirtualNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsForwardingRuleset_ListByVirtualNetwork.json
     */
    /**
     * Sample code: List DNS forwarding rulesets by virtual network.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void listDNSForwardingRulesetsByVirtualNetwork(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .dnsForwardingRulesets()
            .listByVirtualNetwork("sampleResourceGroup", "sampleVirtualNetwork", null, Context.NONE);
    }
}
