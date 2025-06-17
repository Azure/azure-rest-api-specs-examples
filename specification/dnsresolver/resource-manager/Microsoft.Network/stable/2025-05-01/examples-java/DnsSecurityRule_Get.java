
/**
 * Samples for DnsSecurityRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsSecurityRule_Get.json
     */
    /**
     * Sample code: Retrieve DNS security rule for DNS resolver policy.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        retrieveDNSSecurityRuleForDNSResolverPolicy(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsSecurityRules().getWithResponse("sampleResourceGroup", "sampleDnsResolverPolicy",
            "sampleDnsSecurityRule", com.azure.core.util.Context.NONE);
    }
}
