
/**
 * Samples for DnsResolverPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsResolverPolicy_Delete.
     * json
     */
    /**
     * Sample code: Delete DNS resolver policy.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void deleteDNSResolverPolicy(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverPolicies().delete("sampleResourceGroup", "sampleDnsResolverPolicy", null,
            com.azure.core.util.Context.NONE);
    }
}
