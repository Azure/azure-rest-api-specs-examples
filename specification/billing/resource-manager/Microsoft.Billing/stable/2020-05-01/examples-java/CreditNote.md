```java
import com.azure.core.util.Context;

/** Samples for Invoices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CreditNote.json
     */
    /**
     * Sample code: CreditNote.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void creditNote(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().getWithResponse("{billingAccountName}", "{invoiceName}", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-billing_1.0.0-beta.2/sdk/billing/azure-resourcemanager-billing/README.md) on how to add the SDK to your project and authenticate.
