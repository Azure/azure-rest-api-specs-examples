import com.azure.core.util.Context;

/** Samples for MarketplaceAgreements Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/MarketplaceAgreements_Create.json
     */
    /**
     * Sample code: MarketplaceAgreements_Create.
     *
     * @param manager Entry point to ConfluentManager.
     */
    public static void marketplaceAgreementsCreate(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.marketplaceAgreements().createWithResponse(null, Context.NONE);
    }
}
