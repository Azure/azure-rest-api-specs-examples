import com.azure.core.util.Context;

/** Samples for AdaptiveNetworkHardenings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/GetAdaptiveNetworkHardening_example.json
     */
    /**
     * Sample code: Get a single Adaptive Network Hardening resource.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getASingleAdaptiveNetworkHardeningResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .adaptiveNetworkHardenings()
            .getWithResponse("rg1", "Microsoft.Compute", "virtualMachines", "vm1", "default", Context.NONE);
    }
}
