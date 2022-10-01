import com.azure.core.util.Context;

/** Samples for RoleInstances Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/RoleInstanceRestart.json
     */
    /**
     * Sample code: Restart a role instance of a vendor network function.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void restartARoleInstanceOfAVendorNetworkFunction(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.roleInstances().restart("eastus", "testVendor", "testServiceKey", "mrm", Context.NONE);
    }
}
