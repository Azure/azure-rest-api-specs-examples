
/**
 * Samples for Licenses List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Licenses_List.json
     */
    /**
     * Sample code: Licenses_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void licensesList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.licenses().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
