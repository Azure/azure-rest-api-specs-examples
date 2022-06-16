import com.azure.core.util.Context;
import com.azure.resourcemanager.marketplaceordering.models.OfferType;

/** Samples for MarketplaceAgreements Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/GetMarketplaceTerms.json
     */
    /**
     * Sample code: GetMarketplaceTerms.
     *
     * @param manager Entry point to MarketplaceOrderingManager.
     */
    public static void getMarketplaceTerms(
        com.azure.resourcemanager.marketplaceordering.MarketplaceOrderingManager manager) {
        manager
            .marketplaceAgreements()
            .getWithResponse(OfferType.VIRTUALMACHINE, "pubid", "offid", "planid", Context.NONE);
    }
}
