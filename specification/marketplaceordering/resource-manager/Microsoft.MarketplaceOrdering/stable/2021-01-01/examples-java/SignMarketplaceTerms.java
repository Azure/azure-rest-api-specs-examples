import com.azure.core.util.Context;

/** Samples for MarketplaceAgreements Sign. */
public final class Main {
    /*
     * x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/SignMarketplaceTerms.json
     */
    /**
     * Sample code: SetMarketplaceTerms.
     *
     * @param manager Entry point to MarketplaceOrderingManager.
     */
    public static void setMarketplaceTerms(
        com.azure.resourcemanager.marketplaceordering.MarketplaceOrderingManager manager) {
        manager.marketplaceAgreements().signWithResponse("pubid", "offid", "planid", Context.NONE);
    }
}
