import com.azure.core.util.Context;

/** Samples for VirtualNetworkLinks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualNetworkLink_Get.json
     */
    /**
     * Sample code: Retrieve virtual network link to a DNS forwarding ruleset.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void retrieveVirtualNetworkLinkToADNSForwardingRuleset(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .virtualNetworkLinks()
            .getWithResponse(
                "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleVirtualNetworkLink", Context.NONE);
    }
}
