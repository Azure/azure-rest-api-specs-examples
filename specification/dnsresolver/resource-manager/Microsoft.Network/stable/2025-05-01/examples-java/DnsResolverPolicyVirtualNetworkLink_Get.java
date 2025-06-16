
/**
 * Samples for DnsResolverPolicyVirtualNetworkLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * DnsResolverPolicyVirtualNetworkLink_Get.json
     */
    /**
     * Sample code: Retrieve DNS resolver policy virtual network link.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        retrieveDNSResolverPolicyVirtualNetworkLink(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverPolicyVirtualNetworkLinks().getWithResponse("sampleResourceGroup", "sampleDnsResolverPolicy",
            "sampleVirtualNetworkLink", com.azure.core.util.Context.NONE);
    }
}
