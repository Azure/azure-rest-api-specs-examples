
/**
 * Samples for Entities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entities/GetDnsEntityById.json
     */
    /**
     * Sample code: Get a dns entity.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getADnsEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "f4e74920-f2c0-4412-a45f-66d94fdf01f8",
            com.azure.core.util.Context.NONE);
    }
}
