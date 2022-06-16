import com.azure.core.util.Context;

/** Samples for DnsForwardingRulesets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsForwardingRuleset_Delete.json
     */
    /**
     * Sample code: Delete DNS forwarding ruleset.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void deleteDNSForwardingRuleset(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .dnsForwardingRulesets()
            .delete("sampleResourceGroup", "samplednsForwardingRulesetName", null, Context.NONE);
    }
}
