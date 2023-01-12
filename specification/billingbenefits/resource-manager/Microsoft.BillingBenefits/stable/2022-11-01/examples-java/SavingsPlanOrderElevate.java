/** Samples for SavingsPlanOrder Elevate. */
public final class Main {
    /*
     * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanOrderElevate.json
     */
    /**
     * Sample code: SavingsPlanOrderElevate.
     *
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void savingsPlanOrderElevate(
        com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager
            .savingsPlanOrders()
            .elevateWithResponse("20000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
