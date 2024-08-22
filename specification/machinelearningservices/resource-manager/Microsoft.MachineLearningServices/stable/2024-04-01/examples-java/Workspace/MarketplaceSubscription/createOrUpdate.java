
import com.azure.resourcemanager.machinelearning.models.MarketplaceSubscriptionProperties;

/**
 * Samples for MarketplaceSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/MarketplaceSubscription/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Marketplace Subscription.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateWorkspaceMarketplaceSubscription(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.marketplaceSubscriptions().define("string").withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(new MarketplaceSubscriptionProperties().withModelId("string")).create();
    }
}
