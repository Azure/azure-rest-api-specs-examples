Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PriceSheet GetByBillingPeriod. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/PriceSheetForBillingPeriod.json
     */
    /**
     * Sample code: PriceSheetForBillingPeriod.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void priceSheetForBillingPeriod(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.priceSheets().getByBillingPeriodWithResponse("201801", null, null, null, Context.NONE);
    }
}
```
