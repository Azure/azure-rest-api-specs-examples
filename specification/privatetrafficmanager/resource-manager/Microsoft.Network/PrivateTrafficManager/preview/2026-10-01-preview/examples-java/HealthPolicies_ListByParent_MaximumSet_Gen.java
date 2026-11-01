
/**
 * Samples for HealthPolicies ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/HealthPolicies_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: HealthPolicies_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void healthPoliciesListByParentMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.healthPolicies().listByParent("rgprivateTrafficManager", "myProfile", com.azure.core.util.Context.NONE);
    }
}
