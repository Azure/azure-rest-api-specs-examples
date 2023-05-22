/** Samples for Policies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2022-05-01/examples/WafPolicyDelete.json
     */
    /**
     * Sample code: Delete protection policy.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void deleteProtectionPolicy(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.policies().delete("rg1", "Policy1", com.azure.core.util.Context.NONE);
    }
}
