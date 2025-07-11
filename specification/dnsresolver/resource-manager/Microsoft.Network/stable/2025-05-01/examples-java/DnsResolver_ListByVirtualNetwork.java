
/**
 * Samples for DnsResolvers ListByVirtualNetwork.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * DnsResolver_ListByVirtualNetwork.json
     */
    /**
     * Sample code: List DNS resolvers by virtual network.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSResolversByVirtualNetwork(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolvers().listByVirtualNetwork("sampleResourceGroup", "sampleVirtualNetwork", null,
            com.azure.core.util.Context.NONE);
    }
}
