import com.azure.core.util.Context;

/** Samples for OutboundEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/OutboundEndpoint_List.json
     */
    /**
     * Sample code: List outbound endpoints by DNS resolver.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void listOutboundEndpointsByDNSResolver(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.outboundEndpoints().list("sampleResourceGroup", "sampleDnsResolver", null, Context.NONE);
    }
}
