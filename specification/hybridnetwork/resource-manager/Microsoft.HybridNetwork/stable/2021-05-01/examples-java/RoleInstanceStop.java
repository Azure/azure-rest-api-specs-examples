import com.azure.core.util.Context;

/** Samples for RoleInstances Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/RoleInstanceStop.json
     */
    /**
     * Sample code: Stop a role instance of a vendor network function.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void stopARoleInstanceOfAVendorNetworkFunction(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.roleInstances().stop("eastus", "testVendor", "testServiceKey", "mrm", Context.NONE);
    }
}
