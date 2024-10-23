
/**
 * Samples for OutboundEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/
     * OutboundEndpoint_Get.json
     */
    /**
     * Sample code: Retrieve outbound endpoint for DNS resolver.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        retrieveOutboundEndpointForDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.outboundEndpoints().getWithResponse("sampleResourceGroup", "sampleDnsResolver",
            "sampleOutboundEndpoint", com.azure.core.util.Context.NONE);
    }
}
