/** Samples for SavingsPlan Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanItemGet.json
     */
    /**
     * Sample code: SavingsPlanItemGet.
     *
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void savingsPlanItemGet(com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager
            .savingsPlans()
            .getWithResponse(
                "20000000-0000-0000-0000-000000000000",
                "30000000-0000-0000-0000-000000000000",
                null,
                com.azure.core.util.Context.NONE);
    }
}
