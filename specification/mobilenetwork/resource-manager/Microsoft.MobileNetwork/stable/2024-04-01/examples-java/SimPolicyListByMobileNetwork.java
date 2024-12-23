
/**
 * Samples for SimPolicies ListByMobileNetwork.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * SimPolicyListByMobileNetwork.json
     */
    /**
     * Sample code: List SIM policies in a mobile network.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void
        listSIMPoliciesInAMobileNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simPolicies().listByMobileNetwork("testResourceGroupName", "testMobileNetwork",
            com.azure.core.util.Context.NONE);
    }
}
