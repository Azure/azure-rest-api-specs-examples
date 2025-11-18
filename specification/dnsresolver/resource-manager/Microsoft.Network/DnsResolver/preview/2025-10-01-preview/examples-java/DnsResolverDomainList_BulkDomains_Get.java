
/**
 * Samples for DnsResolverDomainLists GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverDomainList_BulkDomains_Get.json
     */
    /**
     * Sample code: Retrieve DNS resolver domain list with bulk number of domains.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void retrieveDNSResolverDomainListWithBulkNumberOfDomains(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverDomainLists().getByResourceGroupWithResponse("sampleResourceGroup",
            "sampleDnsResolverDomainList", com.azure.core.util.Context.NONE);
    }
}
