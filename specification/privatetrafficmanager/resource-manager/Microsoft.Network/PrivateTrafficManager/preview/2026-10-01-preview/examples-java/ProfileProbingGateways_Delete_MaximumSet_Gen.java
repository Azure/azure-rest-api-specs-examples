
/**
 * Samples for ProfileProbingGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ProfileProbingGateways_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ProfileProbingGateways_Delete_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void profileProbingGatewaysDeleteMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.profileProbingGateways().delete("rgprivateTrafficManager", "myProfile", "myProbingGateway",
            com.azure.core.util.Context.NONE);
    }
}
