
/**
 * Samples for MarketplaceAgreements Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/MarketplaceAgreements_Create_MinimumSet_Gen.json
     */
    /**
     * Sample code: Create Confluent Marketplace agreement in the subscription. (MinimumSet).
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void createConfluentMarketplaceAgreementInTheSubscriptionMinimumSet(
        com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.marketplaceAgreements().createWithResponse(null, com.azure.core.util.Context.NONE);
    }
}
