
/**
 * Samples for Authorizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Authorizations_Delete.json
     */
    /**
     * Sample code: Authorizations_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void authorizationsDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.authorizations().delete("group1", "cloud1", "authorization1", com.azure.core.util.Context.NONE);
    }
}
