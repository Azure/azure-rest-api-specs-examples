
/**
 * Samples for DnsResolverPolicyVirtualNetworkLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverPolicyVirtualNetworkLink_Get.json
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
