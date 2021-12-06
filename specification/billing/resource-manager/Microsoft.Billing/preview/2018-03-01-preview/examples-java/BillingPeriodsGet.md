Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-billing_1.0.0-beta.2/sdk/billing/azure-resourcemanager-billing/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for BillingPeriods Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/BillingPeriodsGet.json
     */
    /**
     * Sample code: BillingPeriodsGet.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingPeriodsGet(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingPeriods().getWithResponse("201702-1", Context.NONE);
    }
}
```
