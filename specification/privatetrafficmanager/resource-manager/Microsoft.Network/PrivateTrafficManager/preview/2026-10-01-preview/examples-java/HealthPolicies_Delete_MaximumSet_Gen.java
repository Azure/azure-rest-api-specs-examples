
/**
 * Samples for HealthPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/HealthPolicies_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: HealthPolicies_Delete_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void healthPoliciesDeleteMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.healthPolicies().delete("rgprivateTrafficManager", "myProfile", "myHealthPolicy",
            com.azure.core.util.Context.NONE);
    }
}
