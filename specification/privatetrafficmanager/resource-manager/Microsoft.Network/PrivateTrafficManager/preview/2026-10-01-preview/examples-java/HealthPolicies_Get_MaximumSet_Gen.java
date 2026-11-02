
/**
 * Samples for HealthPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/HealthPolicies_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: HealthPolicies_Get_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void healthPoliciesGetMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.healthPolicies().getWithResponse("rgprivateTrafficManager", "myProfile", "myHealthPolicy",
            com.azure.core.util.Context.NONE);
    }
}
