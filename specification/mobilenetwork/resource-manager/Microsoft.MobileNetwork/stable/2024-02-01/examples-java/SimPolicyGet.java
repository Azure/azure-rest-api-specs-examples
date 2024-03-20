
/**
 * Samples for SimPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimPolicyGet.json
     */
    /**
     * Sample code: Get SIM policy.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getSIMPolicy(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simPolicies().getWithResponse("rg1", "testMobileNetwork", "testPolicy",
            com.azure.core.util.Context.NONE);
    }
}
