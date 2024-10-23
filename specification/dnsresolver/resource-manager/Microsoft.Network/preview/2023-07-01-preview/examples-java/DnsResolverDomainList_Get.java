
/**
 * Samples for DnsResolverDomainLists GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/
     * DnsResolverDomainList_Get.json
     */
    /**
     * Sample code: Retrieve DNS resolver domain list.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void retrieveDNSResolverDomainList(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverDomainLists().getByResourceGroupWithResponse("sampleResourceGroup",
            "sampleDnsResolverDomainList", com.azure.core.util.Context.NONE);
    }
}
