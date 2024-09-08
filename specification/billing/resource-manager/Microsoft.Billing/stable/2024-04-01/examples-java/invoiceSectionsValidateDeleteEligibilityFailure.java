
/**
 * Samples for InvoiceSections ValidateDeleteEligibility.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoiceSectionsValidateDeleteEligibilityFailure.json
     */
    /**
     * Sample code: InvoiceSectionsValidateDeleteEligibilityFailure.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        invoiceSectionsValidateDeleteEligibilityFailure(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoiceSections().validateDeleteEligibilityWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "yyyy-yyyy-yyy-yyy", com.azure.core.util.Context.NONE);
    }
}
