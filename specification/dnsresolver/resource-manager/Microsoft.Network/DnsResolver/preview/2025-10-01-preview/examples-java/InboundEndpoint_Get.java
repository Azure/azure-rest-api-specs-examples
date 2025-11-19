
/**
 * Samples for InboundEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/InboundEndpoint_Get.json
     */
    /**
     * Sample code: Retrieve inbound endpoint for DNS resolver.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        retrieveInboundEndpointForDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.inboundEndpoints().getWithResponse("sampleResourceGroup", "sampleDnsResolver", "sampleInboundEndpoint",
            com.azure.core.util.Context.NONE);
    }
}
