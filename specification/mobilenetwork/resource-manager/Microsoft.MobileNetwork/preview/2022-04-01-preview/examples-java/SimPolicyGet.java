import com.azure.core.util.Context;

/** Samples for SimPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimPolicyGet.json
     */
    /**
     * Sample code: Get SIM policy.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getSIMPolicy(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simPolicies().getWithResponse("rg1", "testMobileNetwork", "testPolicy", Context.NONE);
    }
}
