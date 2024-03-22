
/**
 * Samples for MarketplaceAgreements Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/
     * MarketplaceAgreements_Create.json
     */
    /**
     * Sample code: MarketplaceAgreements_Create.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void marketplaceAgreementsCreate(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.marketplaceAgreements().createWithResponse(null, com.azure.core.util.Context.NONE);
    }
}
