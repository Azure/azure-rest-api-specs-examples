
/**
 * Samples for VirtualNetworkLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/VirtualNetworkLink_Get.json
     */
    /**
     * Sample code: Retrieve virtual network link to a DNS forwarding ruleset.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void retrieveVirtualNetworkLinkToADNSForwardingRuleset(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.virtualNetworkLinks().getWithResponse("sampleResourceGroup", "sampleDnsForwardingRuleset",
            "sampleVirtualNetworkLink", com.azure.core.util.Context.NONE);
    }
}
