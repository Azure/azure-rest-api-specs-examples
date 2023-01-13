/** Samples for SavingsPlan List. */
public final class Main {
    /*
     * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlansInOrderList.json
     */
    /**
     * Sample code: SavingsPlansInOrderList.
     *
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void savingsPlansInOrderList(
        com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager.savingsPlans().list("20000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
