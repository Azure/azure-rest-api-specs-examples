
/**
 * Samples for DnsResolvers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/
     * DnsResolver_ListByResourceGroup.json
     */
    /**
     * Sample code: List DNS resolvers by resource group.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSResolversByResourceGroup(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolvers().listByResourceGroup("sampleResourceGroup", null, com.azure.core.util.Context.NONE);
    }
}
