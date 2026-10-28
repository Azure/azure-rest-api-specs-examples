
/**
 * Samples for AfdOriginGroups ListByProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDOriginGroups_ListByProfile.json
     */
    /**
     * Sample code: AFDOriginGroups_ListByProfile.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDOriginGroupsListByProfile(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdOriginGroups().listByProfile("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
