
/**
 * Samples for OutboundEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/OutboundEndpoint_List.json
     */
    /**
     * Sample code: List outbound endpoints by DNS resolver.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listOutboundEndpointsByDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.outboundEndpoints().list("sampleResourceGroup", "sampleDnsResolver", null,
            com.azure.core.util.Context.NONE);
    }
}
