import com.azure.core.util.Context;

/** Samples for VirtualNetworkLinks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualNetworkLink_Delete.json
     */
    /**
     * Sample code: Delete virtual network link to a DNS forwarding ruleset.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void deleteVirtualNetworkLinkToADNSForwardingRuleset(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .virtualNetworkLinks()
            .delete(
                "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleVirtualNetworkLink", null, Context.NONE);
    }
}
