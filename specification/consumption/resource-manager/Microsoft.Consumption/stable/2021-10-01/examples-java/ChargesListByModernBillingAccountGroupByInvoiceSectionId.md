Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Charges List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ChargesListByModernBillingAccountGroupByInvoiceSectionId.json
     */
    /**
     * Sample code: ChargesListByBillingAccountGroupByInvoiceSectionId-Modern.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void chargesListByBillingAccountGroupByInvoiceSectionIdModern(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .charges()
            .listWithResponse(
                "providers/Microsoft.Billing/billingAccounts/1234:56789/billingProfiles/42425",
                "2019-09-01",
                "2019-09-30",
                null,
                "groupby((properties/invoiceSectionId))",
                Context.NONE);
    }
}
```
