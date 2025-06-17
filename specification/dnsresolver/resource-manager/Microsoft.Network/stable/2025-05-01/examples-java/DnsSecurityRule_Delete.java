
/**
 * Samples for DnsSecurityRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsSecurityRule_Delete.
     * json
     */
    /**
     * Sample code: Delete DNS security rule for DNS resolver policy.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        deleteDNSSecurityRuleForDNSResolverPolicy(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsSecurityRules().delete("sampleResourceGroup", "sampleDnsDnsResolverPolicy", "sampleDnsSecurityRule",
            null, com.azure.core.util.Context.NONE);
    }
}
