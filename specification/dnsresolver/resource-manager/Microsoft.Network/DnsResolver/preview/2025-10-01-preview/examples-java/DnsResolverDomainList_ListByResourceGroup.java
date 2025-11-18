
/**
 * Samples for DnsResolverDomainLists ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverDomainList_ListByResourceGroup.json
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
