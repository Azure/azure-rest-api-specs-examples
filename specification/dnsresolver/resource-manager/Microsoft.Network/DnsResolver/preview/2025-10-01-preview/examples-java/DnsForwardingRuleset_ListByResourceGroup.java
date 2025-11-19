
/**
 * Samples for DnsForwardingRulesets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsForwardingRuleset_ListByResourceGroup.json
     */
    /**
     * Sample code: List DNS forwarding rulesets by resource group.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSForwardingRulesetsByResourceGroup(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsForwardingRulesets().listByResourceGroup("sampleResourceGroup", null,
            com.azure.core.util.Context.NONE);
    }
}
