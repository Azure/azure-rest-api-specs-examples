
/**
 * Samples for DnsForwardingRulesets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsForwardingRuleset_Delete.json
     */
    /**
     * Sample code: Delete DNS forwarding ruleset.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void deleteDNSForwardingRuleset(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsForwardingRulesets().delete("sampleResourceGroup", "samplednsForwardingRulesetName", null,
            com.azure.core.util.Context.NONE);
    }
}
