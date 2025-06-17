
/**
 * Samples for DnsResolverPolicies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsResolverPolicy_Get.
     * json
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
