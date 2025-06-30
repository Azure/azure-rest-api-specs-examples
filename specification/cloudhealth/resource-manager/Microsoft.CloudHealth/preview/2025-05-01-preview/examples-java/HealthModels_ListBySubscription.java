
/**
 * Samples for HealthModels List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/HealthModels_ListBySubscription.json
     */
    /**
     * Sample code: HealthModels_ListBySubscription.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        healthModelsListBySubscription(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.healthModels().list(com.azure.core.util.Context.NONE);
    }
}
