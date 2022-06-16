import com.azure.core.util.Context;

/** Samples for MarketplaceAgreements List. */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/MarketplaceAgreements_List.json
     */
    /**
     * Sample code: MarketplaceAgreements_List.
     *
     * @param manager Entry point to ConfluentManager.
     */
    public static void marketplaceAgreementsList(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.marketplaceAgreements().list(Context.NONE);
    }
}
