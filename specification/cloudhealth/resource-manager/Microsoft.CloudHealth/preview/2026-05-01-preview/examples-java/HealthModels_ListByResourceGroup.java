
/**
 * Samples for HealthModels ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/HealthModels_ListByResourceGroup.json
     */
    /**
     * Sample code: HealthModels_ListByResourceGroup.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        healthModelsListByResourceGroup(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.healthModels().listByResourceGroup("online-store-rg", com.azure.core.util.Context.NONE);
    }
}
