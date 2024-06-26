import com.azure.core.util.Context;

/** Samples for RoleInstances Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/RoleInstanceStart.json
     */
    /**
     * Sample code: Start a role instance of a vendor network function.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void startARoleInstanceOfAVendorNetworkFunction(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.roleInstances().start("eastus", "testVendor", "testServiceKey", "mrm", Context.NONE);
    }
}
