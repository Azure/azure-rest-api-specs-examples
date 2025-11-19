
/**
 * Samples for DnsResolverDomainLists Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverDomainList_Delete.json
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
