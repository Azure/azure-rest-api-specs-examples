
/**
 * Samples for DnsResolverPolicies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/
     * DnsResolverPolicy_ListByResourceGroup.json
     */
    /**
     * Sample code: List DNS resolver policies by resource group.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSResolverPoliciesByResourceGroup(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverPolicies().listByResourceGroup("sampleResourceGroup", null,
            com.azure.core.util.Context.NONE);
    }
}
