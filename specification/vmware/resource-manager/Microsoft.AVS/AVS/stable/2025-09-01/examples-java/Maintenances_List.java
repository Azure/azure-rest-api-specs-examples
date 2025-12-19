
/**
 * Samples for Maintenances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Maintenances_List.json
     */
    /**
     * Sample code: Maintenances_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void maintenancesList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.maintenances().list("group1", "cloud1", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
