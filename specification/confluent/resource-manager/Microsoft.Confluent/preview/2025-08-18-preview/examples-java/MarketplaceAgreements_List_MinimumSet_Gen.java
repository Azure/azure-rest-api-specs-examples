
/**
 * Samples for MarketplaceAgreements List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/MarketplaceAgreements_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: List Confluent marketplace agreements in the subscription. (Minimumset).
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void listConfluentMarketplaceAgreementsInTheSubscriptionMinimumset(
        com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.marketplaceAgreements().list(com.azure.core.util.Context.NONE);
    }
}
