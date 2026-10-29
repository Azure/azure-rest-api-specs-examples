
/**
 * Samples for Profiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Profiles_List.json
     */
    /**
     * Sample code: Profiles_List.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesList(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().list(com.azure.core.util.Context.NONE);
    }
}
