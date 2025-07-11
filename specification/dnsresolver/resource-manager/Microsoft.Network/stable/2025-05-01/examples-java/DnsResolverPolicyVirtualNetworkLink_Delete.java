
/**
 * Samples for DnsResolverPolicyVirtualNetworkLinks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * DnsResolverPolicyVirtualNetworkLink_Delete.json
     */
    /**
     * Sample code: Delete DNS resolver policy virtual network link.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        deleteDNSResolverPolicyVirtualNetworkLink(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverPolicyVirtualNetworkLinks().delete("sampleResourceGroup", "sampleDnsResolverPolicy",
            "sampleVirtualNetworkLink", null, com.azure.core.util.Context.NONE);
    }
}
