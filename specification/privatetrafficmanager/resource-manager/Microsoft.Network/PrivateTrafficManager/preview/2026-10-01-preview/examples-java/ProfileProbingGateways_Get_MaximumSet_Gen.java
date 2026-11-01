
/**
 * Samples for ProfileProbingGateways Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ProfileProbingGateways_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ProfileProbingGateways_Get_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void profileProbingGatewaysGetMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.profileProbingGateways().getWithResponse("rgprivateTrafficManager", "myProfile", "myProbingGateway",
            com.azure.core.util.Context.NONE);
    }
}
