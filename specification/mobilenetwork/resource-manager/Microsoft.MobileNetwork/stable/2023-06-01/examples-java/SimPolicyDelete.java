/** Samples for SimPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimPolicyDelete.json
     */
    /**
     * Sample code: Delete SIM policy.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteSIMPolicy(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simPolicies().delete("rg1", "testMobileNetwork", "testPolicy", com.azure.core.util.Context.NONE);
    }
}
