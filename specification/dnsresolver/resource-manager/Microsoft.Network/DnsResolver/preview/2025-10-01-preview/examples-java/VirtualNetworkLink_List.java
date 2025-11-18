
/**
 * Samples for VirtualNetworkLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/VirtualNetworkLink_List.json
     */
    /**
     * Sample code: List virtual network links to a DNS forwarding ruleset.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void listVirtualNetworkLinksToADNSForwardingRuleset(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.virtualNetworkLinks().list("sampleResourceGroup", "sampleDnsForwardingRuleset", null,
            com.azure.core.util.Context.NONE);
    }
}
