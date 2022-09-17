import com.azure.core.util.Context;

/** Samples for DnsForwardingRulesets ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DnsForwardingRuleset_ListByResourceGroup.json
     */
    /**
     * Sample code: List DNS forwarding rulesets by resource group.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void listDNSForwardingRulesetsByResourceGroup(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsForwardingRulesets().listByResourceGroup("sampleResourceGroup", null, Context.NONE);
    }
}
