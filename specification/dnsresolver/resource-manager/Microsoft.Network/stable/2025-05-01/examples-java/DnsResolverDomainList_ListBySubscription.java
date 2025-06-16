
/**
 * Samples for DnsResolverDomainLists List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * DnsResolverDomainList_ListBySubscription.json
     */
    /**
     * Sample code: List DNS resolver domain lists by subscription.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSResolverDomainListsBySubscription(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverDomainLists().list(null, com.azure.core.util.Context.NONE);
    }
}
