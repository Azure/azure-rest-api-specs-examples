
/**
 * Samples for Entities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Entities_Get.json
     */
    /**
     * Sample code: Entities_Get.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void entitiesGet(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.entities().getWithResponse("rgopenapi", "myHealthModel", "entity1", com.azure.core.util.Context.NONE);
    }
}
