/** Samples for SavingsPlan ListAll. */
public final class Main {
    /*
     * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlansList.json
     */
    /**
     * Sample code: SavingsPlansList.
     *
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void savingsPlansList(com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager
            .savingsPlans()
            .listAll(
                "(properties/archived eq false)",
                "properties/displayName asc",
                "true",
                50.0F,
                null,
                1.0F,
                com.azure.core.util.Context.NONE);
    }
}
