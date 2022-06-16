import com.azure.core.util.Context;

/** Samples for DnsForwardingRulesets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsForwardingRuleset_ListBySubscription.json
     */
    /**
     * Sample code: List DNS forwarding rulesets by subscription.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void listDNSForwardingRulesetsBySubscription(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsForwardingRulesets().list(null, Context.NONE);
    }
}
