import com.azure.core.util.Context;

/** Samples for ForwardingRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ForwardingRule_Delete.json
     */
    /**
     * Sample code: Delete forwarding rule in a DNS forwarding ruleset.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void deleteForwardingRuleInADNSForwardingRuleset(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .forwardingRules()
            .deleteWithResponse(
                "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleForwardingRule", null, Context.NONE);
    }
}
