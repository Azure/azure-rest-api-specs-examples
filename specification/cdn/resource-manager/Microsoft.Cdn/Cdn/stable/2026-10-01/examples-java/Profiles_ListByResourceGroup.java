
/**
 * Samples for Profiles ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Profiles_ListByResourceGroup.json
     */
    /**
     * Sample code: Profiles_ListByResourceGroup.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesListByResourceGroup(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().listByResourceGroup("RG", com.azure.core.util.Context.NONE);
    }
}
