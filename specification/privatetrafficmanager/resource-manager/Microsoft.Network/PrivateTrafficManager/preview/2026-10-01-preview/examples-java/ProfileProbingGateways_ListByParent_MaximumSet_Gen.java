
/**
 * Samples for ProfileProbingGateways ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ProfileProbingGateways_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: ProfileProbingGateways_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void profileProbingGatewaysListByParentMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.profileProbingGateways().listByParent("rgprivateTrafficManager", "myProfile",
            com.azure.core.util.Context.NONE);
    }
}
