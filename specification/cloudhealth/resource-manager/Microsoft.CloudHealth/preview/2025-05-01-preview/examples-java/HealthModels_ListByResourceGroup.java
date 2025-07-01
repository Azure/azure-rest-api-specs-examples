
/**
 * Samples for HealthModels ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/HealthModels_ListByResourceGroup.json
     */
    /**
     * Sample code: HealthModels_ListByResourceGroup.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        healthModelsListByResourceGroup(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.healthModels().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
