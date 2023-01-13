/** Samples for SavingsPlanOrderAlias Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanOrderAliasGet.json
     */
    /**
     * Sample code: SavingsPlanOrderAliasGet.
     *
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void savingsPlanOrderAliasGet(
        com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager.savingsPlanOrderAlias().getWithResponse("spAlias123", com.azure.core.util.Context.NONE);
    }
}
