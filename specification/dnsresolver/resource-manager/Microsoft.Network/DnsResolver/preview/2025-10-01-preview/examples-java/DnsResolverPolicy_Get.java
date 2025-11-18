
/**
 * Samples for DnsResolverPolicies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverPolicy_Get.json
     */
    /**
     * Sample code: Retrieve DNS resolver policy.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void retrieveDNSResolverPolicy(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverPolicies().getByResourceGroupWithResponse("sampleResourceGroup", "sampleDnsResolverPolicy",
            com.azure.core.util.Context.NONE);
    }
}
