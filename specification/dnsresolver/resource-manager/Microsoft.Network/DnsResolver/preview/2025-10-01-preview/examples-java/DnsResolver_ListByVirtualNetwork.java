
/**
 * Samples for DnsResolvers ListByVirtualNetwork.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolver_ListByVirtualNetwork.json
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
