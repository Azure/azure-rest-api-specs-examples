
/**
 * Samples for DnsForwardingRulesets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/
     * DnsForwardingRuleset_Get.json
     */
    /**
     * Sample code: Retrieve DNS forwarding ruleset.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void retrieveDNSForwardingRuleset(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsForwardingRulesets().getByResourceGroupWithResponse("sampleResourceGroup",
            "sampleDnsForwardingRuleset", com.azure.core.util.Context.NONE);
    }
}
