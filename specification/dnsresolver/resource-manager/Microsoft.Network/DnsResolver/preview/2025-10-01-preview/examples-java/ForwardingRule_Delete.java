
/**
 * Samples for ForwardingRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ForwardingRule_Delete.json
     */
    /**
     * Sample code: Delete forwarding rule in a DNS forwarding ruleset.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        deleteForwardingRuleInADNSForwardingRuleset(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.forwardingRules().deleteWithResponse("sampleResourceGroup", "sampleDnsForwardingRuleset",
            "sampleForwardingRule", null, com.azure.core.util.Context.NONE);
    }
}
