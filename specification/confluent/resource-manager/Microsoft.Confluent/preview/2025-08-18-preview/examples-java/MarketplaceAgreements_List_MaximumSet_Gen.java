
/**
 * Samples for MarketplaceAgreements List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/MarketplaceAgreements_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: List Confluent marketplace agreements in the subscription. (Maximumset).
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void listConfluentMarketplaceAgreementsInTheSubscriptionMaximumset(
        com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.marketplaceAgreements().list(com.azure.core.util.Context.NONE);
    }
}
