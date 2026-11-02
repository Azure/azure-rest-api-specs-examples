
/**
 * Samples for Profiles ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Profiles_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Profiles_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void profilesListByResourceGroupMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.profiles().listByResourceGroup("rgprivateTrafficManager", com.azure.core.util.Context.NONE);
    }
}
