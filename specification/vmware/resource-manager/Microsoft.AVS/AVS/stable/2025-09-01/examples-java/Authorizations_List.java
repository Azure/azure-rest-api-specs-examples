
/**
 * Samples for Authorizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Authorizations_List.json
     */
    /**
     * Sample code: Authorizations_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void authorizationsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.authorizations().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
