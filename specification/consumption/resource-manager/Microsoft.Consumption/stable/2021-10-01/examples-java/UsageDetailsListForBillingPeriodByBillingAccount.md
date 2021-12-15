Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-consumption_1.0.0-beta.3/sdk/consumption/azure-resourcemanager-consumption/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for UsageDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsListForBillingPeriodByBillingAccount.json
     */
    /**
     * Sample code: BillingAccountUsageDetailsListForBillingPeriod-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void billingAccountUsageDetailsListForBillingPeriodLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .usageDetails()
            .list("providers/Microsoft.Billing/BillingAccounts/1234", null, null, null, null, null, Context.NONE);
    }
}
```
