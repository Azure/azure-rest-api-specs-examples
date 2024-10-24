
/**
 * Samples for DnsResolverDomainLists ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/
     * DnsResolverDomainList_ListByResourceGroup.json
     */
    /**
     * Sample code: List DNS resolver domain lists by resource group.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSResolverDomainListsByResourceGroup(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverDomainLists().listByResourceGroup("sampleResourceGroup", null,
            com.azure.core.util.Context.NONE);
    }
}
