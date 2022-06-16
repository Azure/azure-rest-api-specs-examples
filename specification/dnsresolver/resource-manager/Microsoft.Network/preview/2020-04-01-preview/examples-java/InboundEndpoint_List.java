import com.azure.core.util.Context;

/** Samples for InboundEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/InboundEndpoint_List.json
     */
    /**
     * Sample code: List inbound endpoints by DNS resolver.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void listInboundEndpointsByDNSResolver(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.inboundEndpoints().list("sampleResourceGroup", "sampleDnsResolver", null, Context.NONE);
    }
}
