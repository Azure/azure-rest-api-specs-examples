import com.azure.core.util.Context;

/** Samples for DnsResolvers ListByVirtualNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsResolver_ListByVirtualNetwork.json
     */
    /**
     * Sample code: List DNS resolvers by virtual network.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void listDNSResolversByVirtualNetwork(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolvers().listByVirtualNetwork("sampleResourceGroup", "sampleVirtualNetwork", null, Context.NONE);
    }
}
