
/**
 * Samples for InboundEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/InboundEndpoint_List.json
     */
    /**
     * Sample code: List inbound endpoints by DNS resolver.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listInboundEndpointsByDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.inboundEndpoints().list("sampleResourceGroup", "sampleDnsResolver", null,
            com.azure.core.util.Context.NONE);
    }
}
