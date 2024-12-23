
/**
 * Samples for CustomRecommendations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/CustomRecommendations/
     * ListBySubscriptionCustomRecommendations_example.json
     */
    /**
     * Sample code: List custom recommendations by subscription scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        listCustomRecommendationsBySubscriptionScope(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customRecommendations().list("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
            com.azure.core.util.Context.NONE);
    }
}
