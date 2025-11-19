
/**
 * Samples for DnsResolvers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolver_Delete.json
     */
    /**
     * Sample code: Delete DNS resolver.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void deleteDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolvers().delete("sampleResourceGroup", "sampleDnsResolver", null,
            com.azure.core.util.Context.NONE);
    }
}
