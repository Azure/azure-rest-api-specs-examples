
/**
 * Samples for DnsResolvers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolver_Get.json
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
