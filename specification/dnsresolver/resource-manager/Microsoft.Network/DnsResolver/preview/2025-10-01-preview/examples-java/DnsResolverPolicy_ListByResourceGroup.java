
/**
 * Samples for DnsResolverPolicies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverPolicy_ListByResourceGroup.json
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
