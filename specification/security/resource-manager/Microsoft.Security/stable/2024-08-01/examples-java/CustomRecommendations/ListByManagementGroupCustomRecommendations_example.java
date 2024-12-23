
/**
 * Samples for CustomRecommendations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/CustomRecommendations/
     * ListByManagementGroupCustomRecommendations_example.json
     */
    /**
     * Sample code: List custom recommendations by management group scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        listCustomRecommendationsByManagementGroupScope(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customRecommendations().list("providers/Microsoft.Management/managementGroups/contoso",
            com.azure.core.util.Context.NONE);
    }
}
