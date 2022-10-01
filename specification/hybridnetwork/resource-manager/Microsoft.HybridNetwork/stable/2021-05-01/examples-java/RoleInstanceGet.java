import com.azure.core.util.Context;

/** Samples for RoleInstances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/RoleInstanceGet.json
     */
    /**
     * Sample code: Get the operational state of role instance of vendor network function.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getTheOperationalStateOfRoleInstanceOfVendorNetworkFunction(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.roleInstances().getWithResponse("eastus", "testVendor", "testServiceKey", "mrm", Context.NONE);
    }
}
