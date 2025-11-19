
/**
 * Samples for OutboundEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/OutboundEndpoint_Delete.json
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
