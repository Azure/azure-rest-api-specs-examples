
/**
 * Samples for Profiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Profiles_Delete.json
     */
    /**
     * Sample code: Profiles_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().delete("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
