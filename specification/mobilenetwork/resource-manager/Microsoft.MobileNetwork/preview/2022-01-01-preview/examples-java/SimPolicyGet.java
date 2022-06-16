import com.azure.core.util.Context;

/** Samples for SimPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SimPolicyGet.json
     */
    /**
     * Sample code: Get sim policy.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getSimPolicy(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simPolicies().getWithResponse("rg1", "testMobileNetwork", "testPolicy", Context.NONE);
    }
}
