
/**
 * Samples for ForwardingRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/
     * ForwardingRule_List.json
     */
    /**
     * Sample code: List forwarding rules in a DNS forwarding ruleset.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listForwardingRulesInADNSForwardingRuleset(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.forwardingRules().list("sampleResourceGroup", "sampleDnsForwardingRuleset", null,
            com.azure.core.util.Context.NONE);
    }
}
