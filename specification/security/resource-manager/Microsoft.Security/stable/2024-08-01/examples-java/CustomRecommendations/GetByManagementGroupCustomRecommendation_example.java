
/**
 * Samples for CustomRecommendations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/CustomRecommendations/
     * GetByManagementGroupCustomRecommendation_example.json
     */
    /**
     * Sample code: Get a custom recommendation over management group scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        getACustomRecommendationOverManagementGroupScope(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customRecommendations().getWithResponse("providers/Microsoft.Management/managementGroups/contoso",
            "1f3afdf9-d0c9-4c3d-847f-89da613e70a8", com.azure.core.util.Context.NONE);
    }
}
