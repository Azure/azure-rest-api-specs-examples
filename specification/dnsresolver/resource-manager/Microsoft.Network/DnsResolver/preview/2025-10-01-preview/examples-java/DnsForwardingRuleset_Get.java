
/**
 * Samples for DnsForwardingRulesets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsForwardingRuleset_Get.json
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
