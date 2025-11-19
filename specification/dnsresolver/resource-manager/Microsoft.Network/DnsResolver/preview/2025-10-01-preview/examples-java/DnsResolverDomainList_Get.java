
/**
 * Samples for DnsResolverDomainLists GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverDomainList_Get.json
     */
    /**
     * Sample code: Retrieve DNS resolver domain list with less than 1000 domains.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void retrieveDNSResolverDomainListWithLessThan1000Domains(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverDomainLists().getByResourceGroupWithResponse("sampleResourceGroup",
            "sampleDnsResolverDomainList", com.azure.core.util.Context.NONE);
    }
}
