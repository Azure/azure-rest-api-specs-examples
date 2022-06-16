import com.azure.core.util.Context;

/** Samples for Charges List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ChargesForBillingPeriodByDepartment.json
     */
    /**
     * Sample code: ChangesForBillingPeriodByDepartment-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void changesForBillingPeriodByDepartmentLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .charges()
            .listWithResponse(
                "providers/Microsoft.Billing/BillingAccounts/1234/departments/42425",
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
