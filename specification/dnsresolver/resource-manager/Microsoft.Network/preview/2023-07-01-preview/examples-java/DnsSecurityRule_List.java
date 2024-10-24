
/**
 * Samples for DnsSecurityRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/
     * DnsSecurityRule_List.json
     */
    /**
     * Sample code: List DNS security rules by DNS resolver policy.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        listDNSSecurityRulesByDNSResolverPolicy(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsSecurityRules().list("sampleResourceGroup", "sampleDnsResolverPolicy", null,
            com.azure.core.util.Context.NONE);
    }
}
