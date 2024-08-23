
/**
 * Samples for MarketplaceSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/MarketplaceSubscription/delete.json
     */
    /**
     * Sample code: Delete Workspace Marketplace Subscription.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteWorkspaceMarketplaceSubscription(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.marketplaceSubscriptions().delete("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
