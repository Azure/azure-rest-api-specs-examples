
/**
 * Samples for Profiles ListResourceUsage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Profiles_ListResourceUsage.json
     */
    /**
     * Sample code: Profiles_ListResourceUsage.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesListResourceUsage(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().listResourceUsage("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
