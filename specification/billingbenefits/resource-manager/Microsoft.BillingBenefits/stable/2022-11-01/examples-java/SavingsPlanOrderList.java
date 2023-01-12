/** Samples for SavingsPlanOrder List. */
public final class Main {
    /*
     * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanOrderList.json
     */
    /**
     * Sample code: SavingsPlanOrderList.
     *
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void savingsPlanOrderList(com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager.savingsPlanOrders().list(com.azure.core.util.Context.NONE);
    }
}
