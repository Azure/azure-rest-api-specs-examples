import com.azure.core.util.Context;

/** Samples for DnsResolvers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DnsResolver_ListByResourceGroup.json
     */
    /**
     * Sample code: List DNS resolvers by resource group.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void listDNSResolversByResourceGroup(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolvers().listByResourceGroup("sampleResourceGroup", null, Context.NONE);
    }
}
