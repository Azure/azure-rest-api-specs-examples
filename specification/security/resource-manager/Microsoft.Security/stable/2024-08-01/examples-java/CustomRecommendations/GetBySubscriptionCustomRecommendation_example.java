
/**
 * Samples for CustomRecommendations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/CustomRecommendations/
     * GetBySubscriptionCustomRecommendation_example.json
     */
    /**
     * Sample code: Get a custom recommendation over subscription scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        getACustomRecommendationOverSubscriptionScope(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customRecommendations().getWithResponse("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
            "1f3afdf9-d0c9-4c3d-847f-89da613e70a8", com.azure.core.util.Context.NONE);
    }
}
