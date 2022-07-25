import com.azure.core.util.Context;

/** Samples for SimPolicies ListByMobileNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimPolicyListByMobileNetwork.json
     */
    /**
     * Sample code: List SIM policies in a mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSIMPoliciesInAMobileNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simPolicies().listByMobileNetwork("testResourceGroupName", "testMobileNetwork", Context.NONE);
    }
}
