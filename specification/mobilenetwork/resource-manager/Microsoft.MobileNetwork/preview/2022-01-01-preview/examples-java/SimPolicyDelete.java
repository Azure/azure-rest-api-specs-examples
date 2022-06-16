import com.azure.core.util.Context;

/** Samples for SimPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SimPolicyDelete.json
     */
    /**
     * Sample code: Delete sim policy.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteSimPolicy(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simPolicies().delete("rg1", "testMobileNetwork", "testPolicy", Context.NONE);
    }
}
