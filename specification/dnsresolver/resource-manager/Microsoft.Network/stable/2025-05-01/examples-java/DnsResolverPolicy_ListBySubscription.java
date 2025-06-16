
/**
 * Samples for DnsResolverPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * DnsResolverPolicy_ListBySubscription.json
     */
    /**
     * Sample code: List DNS resolver policies by subscription.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSResolverPoliciesBySubscription(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverPolicies().list(null, com.azure.core.util.Context.NONE);
    }
}
