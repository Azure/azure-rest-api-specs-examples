
/**
 * Samples for Maintenances Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Maintenances_Get.json
     */
    /**
     * Sample code: Maintenances_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void maintenancesGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.maintenances().getWithResponse("group1", "cloud1", "maintenance1", com.azure.core.util.Context.NONE);
    }
}
