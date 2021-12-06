Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-billing_1.0.0-beta.2/sdk/billing/azure-resourcemanager-billing/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for InvoiceSections ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionsListByBillingProfile.json
     */
    /**
     * Sample code: InvoiceSectionsListByBillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSectionsListByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoiceSections().listByBillingProfile("{billingAccountName}", "{billingProfileName}", Context.NONE);
    }
}
```
