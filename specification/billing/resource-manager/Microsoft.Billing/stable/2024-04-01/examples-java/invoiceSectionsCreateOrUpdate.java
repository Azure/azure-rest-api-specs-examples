
import com.azure.resourcemanager.billing.fluent.models.InvoiceSectionInner;
import com.azure.resourcemanager.billing.models.InvoiceSectionProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for InvoiceSections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoiceSectionsCreateOrUpdate
     * .json
     */
    /**
     * Sample code: InvoiceSectionsCreateOrUpdate.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSectionsCreateOrUpdate(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoiceSections().createOrUpdate(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "invoice-section-1",
            new InvoiceSectionInner().withProperties(new InvoiceSectionProperties().withDisplayName("Invoice Section 1")
                .withTags(mapOf("costCategory", "Support", "pcCode", "fakeTokenPlaceholder"))),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
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
