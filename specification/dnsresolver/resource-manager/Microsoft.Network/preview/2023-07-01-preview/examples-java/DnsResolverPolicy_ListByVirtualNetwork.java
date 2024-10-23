
/**
 * Samples for DnsResolverPolicies ListByVirtualNetwork.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/
     * DnsResolverPolicy_ListByVirtualNetwork.json
     */
    /**
     * Sample code: List DNS resolver policies by virtual network.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSResolverPoliciesByVirtualNetwork(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverPolicies().listByVirtualNetwork("sampleResourceGroup", "sampleVirtualNetwork",
            com.azure.core.util.Context.NONE);
    }
}
