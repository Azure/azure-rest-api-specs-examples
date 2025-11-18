
/**
 * Samples for DnsResolverPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverPolicy_ListBySubscription.json
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
