
/**
 * Samples for Relationships Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Relationships_Get.json
     */
    /**
     * Sample code: Relationships_Get.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void relationshipsGet(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.relationships().getWithResponse("rgopenapi", "myHealthModel", "Ue-21-F3M12V3w-13x18F8H-7HOk--kq6tP-HB",
            com.azure.core.util.Context.NONE);
    }
}
