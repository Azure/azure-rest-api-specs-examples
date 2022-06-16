import com.azure.core.util.Context;

/** Samples for InboundEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/InboundEndpoint_Delete.json
     */
    /**
     * Sample code: Delete inbound endpoint for DNS resolver.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void deleteInboundEndpointForDNSResolver(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .inboundEndpoints()
            .delete("sampleResourceGroup", "sampleDnsResolver", "sampleInboundEndpoint", null, Context.NONE);
    }
}
