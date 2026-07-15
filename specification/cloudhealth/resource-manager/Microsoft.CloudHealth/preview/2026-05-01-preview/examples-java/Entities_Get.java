
/**
 * Samples for Entities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Entities_Get.json
     */
    /**
     * Sample code: Entities_Get.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void entitiesGet(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.entities().getWithResponse("online-store-rg", "online-store", "orders-db",
            com.azure.core.util.Context.NONE);
    }
}
