
/**
 * Samples for HealthModels Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/HealthModels_Delete.json
     */
    /**
     * Sample code: HealthModels_Delete.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void healthModelsDelete(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.healthModels().delete("rgopenapi", "model1", com.azure.core.util.Context.NONE);
    }
}
