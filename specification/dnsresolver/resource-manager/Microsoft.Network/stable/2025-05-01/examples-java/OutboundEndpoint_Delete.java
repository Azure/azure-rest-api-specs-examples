
/**
 * Samples for OutboundEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/OutboundEndpoint_Delete.
     * json
     */
    /**
     * Sample code: Delete outbound endpoint for DNS resolver.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        deleteOutboundEndpointForDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.outboundEndpoints().delete("sampleResourceGroup", "sampleDnsResolver", "sampleOutboundEndpoint", null,
            com.azure.core.util.Context.NONE);
    }
}
