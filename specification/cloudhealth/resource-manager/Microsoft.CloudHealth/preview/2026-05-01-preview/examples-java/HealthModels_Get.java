
/**
 * Samples for HealthModels GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/HealthModels_Get.json
     */
    /**
     * Sample code: HealthModels_Get.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void healthModelsGet(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.healthModels().getByResourceGroupWithResponse("online-store-rg", "online-store",
            com.azure.core.util.Context.NONE);
    }
}
