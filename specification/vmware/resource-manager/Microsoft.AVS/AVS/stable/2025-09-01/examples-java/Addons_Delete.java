
/**
 * Samples for Addons Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Addons_Delete.json
     */
    /**
     * Sample code: Addons_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void addonsDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().delete("group1", "cloud1", "srm", com.azure.core.util.Context.NONE);
    }
}
