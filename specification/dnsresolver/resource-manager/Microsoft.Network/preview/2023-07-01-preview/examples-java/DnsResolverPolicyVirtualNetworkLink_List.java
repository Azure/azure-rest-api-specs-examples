
/**
 * Samples for DnsResolverPolicyVirtualNetworkLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/
     * DnsResolverPolicyVirtualNetworkLink_List.json
     */
    /**
     * Sample code: List DNS resolver policy virtual network links by DNS resolver policy.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void listDNSResolverPolicyVirtualNetworkLinksByDNSResolverPolicy(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverPolicyVirtualNetworkLinks().list("sampleResourceGroup", "sampleDnsResolverPolicy", null,
            com.azure.core.util.Context.NONE);
    }
}
