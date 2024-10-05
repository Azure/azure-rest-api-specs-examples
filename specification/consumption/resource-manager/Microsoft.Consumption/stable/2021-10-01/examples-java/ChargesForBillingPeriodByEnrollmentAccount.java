
/**
 * Samples for Charges List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ChargesForBillingPeriodByEnrollmentAccount.json
     */
    /**
     * Sample code: ChangesForBillingPeriodByEnrollmentAccount-Legacy.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void changesForBillingPeriodByEnrollmentAccountLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.charges().listWithResponse("providers/Microsoft.Billing/BillingAccounts/1234/enrollmentAccounts/42425",
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
