
/**
 * Samples for CustomRecommendations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/CustomRecommendations/
     * DeleteBySecurityConnectorCustomRecommendation_example.json
     */
    /**
     * Sample code: Delete a custom recommendation over security connector scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteACustomRecommendationOverSecurityConnectorScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customRecommendations().deleteByResourceGroupWithResponse(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector",
            "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", com.azure.core.util.Context.NONE);
    }
}
