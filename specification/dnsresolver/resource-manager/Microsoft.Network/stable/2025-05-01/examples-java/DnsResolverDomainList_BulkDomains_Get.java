
/**
 * Samples for DnsResolverDomainLists GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * DnsResolverDomainList_BulkDomains_Get.json
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
