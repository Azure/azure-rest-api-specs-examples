
/**
 * Samples for DnsResolvers List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * DnsResolver_ListBySubscription.json
     */
    /**
     * Sample code: List DNS resolvers by subscription.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSResolversBySubscription(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolvers().list(null, com.azure.core.util.Context.NONE);
    }
}
