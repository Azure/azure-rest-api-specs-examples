
/**
 * Samples for DnsResolverPolicyVirtualNetworkLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverPolicyVirtualNetworkLink_List.json
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
