Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Charges List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ChargesListByModernInvoiceSectionId.json
     */
    /**
     * Sample code: ChargesListByInvoiceSectionId-Modern.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void chargesListByInvoiceSectionIdModern(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .charges()
            .listWithResponse(
                "providers/Microsoft.Billing/BillingAccounts/1234:56789/invoiceSections/97531",
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
```
