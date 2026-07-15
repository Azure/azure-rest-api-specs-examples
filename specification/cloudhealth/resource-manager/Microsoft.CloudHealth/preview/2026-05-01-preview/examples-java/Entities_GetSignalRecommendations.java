
/**
 * Samples for Entities GetSignalRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Entities_GetSignalRecommendations.json
     */
    /**
     * Sample code: Entities_GetSignalRecommendations.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        entitiesGetSignalRecommendations(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.entities().getSignalRecommendationsWithResponse("online-store-rg", "online-store", "orders-db",
            com.azure.core.util.Context.NONE);
    }
}
