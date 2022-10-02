import com.azure.core.util.Context;

/** Samples for RoleInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/RoleInstanceListByVendorNetworkFunction.json
     */
    /**
     * Sample code: List all role instances of vendor network function.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listAllRoleInstancesOfVendorNetworkFunction(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.roleInstances().list("eastus", "testVendor", "testServiceKey", Context.NONE);
    }
}
