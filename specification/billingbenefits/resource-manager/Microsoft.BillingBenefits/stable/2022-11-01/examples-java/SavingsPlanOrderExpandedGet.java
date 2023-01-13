/** Samples for SavingsPlanOrder Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanOrderExpandedGet.json
     */
    /**
     * Sample code: SavingsPlanOrderWithExpandedPaymentsGet.
     *
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void savingsPlanOrderWithExpandedPaymentsGet(
        com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager
            .savingsPlanOrders()
            .getWithResponse("20000000-0000-0000-0000-000000000000", "schedule", com.azure.core.util.Context.NONE);
    }
}
