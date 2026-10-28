
/**
 * Samples for AfdProfiles ListResourceUsage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDProfiles_ListResourceUsage.json
     */
    /**
     * Sample code: AFDProfiles_ListResourceUsage.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDProfilesListResourceUsage(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdProfiles().listResourceUsage("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
