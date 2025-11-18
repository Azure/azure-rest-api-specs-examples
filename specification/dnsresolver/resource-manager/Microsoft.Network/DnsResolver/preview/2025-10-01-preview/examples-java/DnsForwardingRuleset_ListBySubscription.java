
/**
 * Samples for DnsForwardingRulesets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsForwardingRuleset_ListBySubscription.json
     */
    /**
     * Sample code: List DNS forwarding rulesets by subscription.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSForwardingRulesetsBySubscription(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsForwardingRulesets().list(null, com.azure.core.util.Context.NONE);
    }
}
