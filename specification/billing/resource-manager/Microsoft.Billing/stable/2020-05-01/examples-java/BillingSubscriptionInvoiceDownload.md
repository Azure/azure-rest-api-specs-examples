Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-billing_1.0.0-beta.2/sdk/billing/azure-resourcemanager-billing/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Invoices DownloadBillingSubscriptionInvoice. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionInvoiceDownload.json
     */
    /**
     * Sample code: BillingSubscriptionInvoiceDownload.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionInvoiceDownload(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().downloadBillingSubscriptionInvoice("{invoiceName}", "DRS_12345", Context.NONE);
    }
}
```
