
import com.azure.resourcemanager.cloudhealth.models.DiscoveryRuleProperties;
import com.azure.resourcemanager.cloudhealth.models.DiscoveryRuleRecommendedSignalsBehavior;
import com.azure.resourcemanager.cloudhealth.models.DiscoveryRuleRelationshipDiscoveryBehavior;

/**
 * Samples for DiscoveryRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/DiscoveryRules_CreateOrUpdate.json
     */
    /**
     * Sample code: DiscoveryRules_CreateOrUpdate.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void discoveryRulesCreateOrUpdate(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.discoveryRules().define("myDiscoveryRule").withExistingHealthmodel("myResourceGroup", "myHealthModel")
            .withProperties(new DiscoveryRuleProperties().withDisplayName("myDisplayName").withResourceGraphQuery(
                "resources | where subscriptionId == '7ddfffd7-9b32-40df-1234-828cbd55d6f4' | where resourceGroup == 'my-rg'")
                .withAuthenticationSetting("authSetting1")
                .withDiscoverRelationships(DiscoveryRuleRelationshipDiscoveryBehavior.ENABLED)
                .withAddRecommendedSignals(DiscoveryRuleRecommendedSignalsBehavior.ENABLED))
            .create();
    }
}
