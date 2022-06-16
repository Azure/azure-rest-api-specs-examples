import com.azure.core.util.Context;

/** Samples for DnsResolvers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsResolver_Get.json
     */
    /**
     * Sample code: Retrieve DNS resolver.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void retrieveDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolvers().getByResourceGroupWithResponse("sampleResourceGroup", "sampleDnsResolver", Context.NONE);
    }
}
