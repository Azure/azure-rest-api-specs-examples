
/**
 * Samples for DnsResolvers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsResolver_Get.json
     */
    /**
     * Sample code: Retrieve DNS resolver.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void retrieveDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolvers().getByResourceGroupWithResponse("sampleResourceGroup", "sampleDnsResolver",
            com.azure.core.util.Context.NONE);
    }
}
