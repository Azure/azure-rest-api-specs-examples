
/**
 * Samples for MarketplaceSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/MarketplaceSubscription/get.json
     */
    /**
     * Sample code: Get Workspace Marketplace Subscription.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceMarketplaceSubscription(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.marketplaceSubscriptions().getWithResponse("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
