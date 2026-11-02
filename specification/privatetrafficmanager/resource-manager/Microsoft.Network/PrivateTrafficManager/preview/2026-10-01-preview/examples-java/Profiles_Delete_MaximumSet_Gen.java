
/**
 * Samples for Profiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Profiles_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Profiles_Delete_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void
        profilesDeleteMaximumSet(com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.profiles().delete("rgprivateTrafficManager", "myProfile", com.azure.core.util.Context.NONE);
    }
}
