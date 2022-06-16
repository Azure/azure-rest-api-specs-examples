import com.azure.core.util.Context;

/** Samples for SimPolicies ListByMobileNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyListByMobileNetwork.json
     */
    /**
     * Sample code: List sim policies in a mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSimPoliciesInAMobileNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simPolicies().listByMobileNetwork("testResourceGroupName", "testMobileNetwork", Context.NONE);
    }
}
