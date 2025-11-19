
/**
 * Samples for DnsSecurityRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsSecurityRule_Delete.json
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
