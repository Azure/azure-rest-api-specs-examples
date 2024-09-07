
/**
 * Samples for InvoiceSections ValidateDeleteEligibility.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoiceSectionsValidateDeleteEligibilitySuccess.json
     */
    /**
     * Sample code: InvoiceSectionsValidateDeleteEligibilitySuccess.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        invoiceSectionsValidateDeleteEligibilitySuccess(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoiceSections().validateDeleteEligibilityWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "yyyy-yyyy-yyy-yyy", com.azure.core.util.Context.NONE);
    }
}
