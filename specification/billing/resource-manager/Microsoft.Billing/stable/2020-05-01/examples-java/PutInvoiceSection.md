Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-billing_1.0.0-beta.2/sdk/billing/azure-resourcemanager-billing/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.billing.fluent.models.InvoiceSectionInner;
import java.util.HashMap;
import java.util.Map;

/** Samples for InvoiceSections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutInvoiceSection.json
     */
    /**
     * Sample code: PutInvoiceSection.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void putInvoiceSection(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .invoiceSections()
            .createOrUpdate(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                new InvoiceSectionInner()
                    .withDisplayName("invoiceSection1")
                    .withLabels(mapOf("costCategory", "Support", "pcCode", "A123456")),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```
