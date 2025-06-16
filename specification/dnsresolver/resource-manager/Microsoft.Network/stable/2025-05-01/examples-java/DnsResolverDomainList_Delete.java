
/**
 * Samples for DnsResolverDomainLists Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * DnsResolverDomainList_Delete.json
     */
    /**
     * Sample code: Delete DNS resolver domain list.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void deleteDNSResolverDomainList(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverDomainLists().delete("sampleResourceGroup", "sampleDnsResolverDomainList", null,
            com.azure.core.util.Context.NONE);
    }
}
