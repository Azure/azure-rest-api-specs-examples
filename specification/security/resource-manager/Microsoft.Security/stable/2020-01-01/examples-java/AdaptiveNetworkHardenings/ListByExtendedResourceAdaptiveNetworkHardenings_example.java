
/**
 * Samples for AdaptiveNetworkHardenings ListByExtendedResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/
     * ListByExtendedResourceAdaptiveNetworkHardenings_example.json
     */
    /**
     * Sample code: List Adaptive Network Hardenings resources of an extended resource.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listAdaptiveNetworkHardeningsResourcesOfAnExtendedResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.adaptiveNetworkHardenings().listByExtendedResource("rg1", "Microsoft.Compute", "virtualMachines", "vm1",
            com.azure.core.util.Context.NONE);
    }
}
