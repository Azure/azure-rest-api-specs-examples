
/**
 * Samples for MarketplaceSubscriptions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/MarketplaceSubscription/list.json
     */
    /**
     * Sample code: List Workspace Marketplace Subscription.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceMarketplaceSubscription(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.marketplaceSubscriptions().list("test-rg", "my-aml-workspace", null, com.azure.core.util.Context.NONE);
    }
}
