Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-marketplaceordering_1.0.0-beta.2/sdk/marketplaceordering/azure-resourcemanager-marketplaceordering/README.md) on how to add the SDK to your project and authenticate.

```java
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
```
