
/**
 * Samples for InboundEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/InboundEndpoint_Delete.json
     */
    /**
     * Sample code: Delete inbound endpoint for DNS resolver.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        deleteInboundEndpointForDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.inboundEndpoints().delete("sampleResourceGroup", "sampleDnsResolver", "sampleInboundEndpoint", null,
            com.azure.core.util.Context.NONE);
    }
}
