import com.azure.core.util.Context;

/** Samples for Charges List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ChargesListForDepartmentFilterByStartEndDate.json
     */
    /**
     * Sample code: ChargesListByDepartment-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void chargesListByDepartmentLegacy(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .charges()
            .listWithResponse(
                "providers/Microsoft.Billing/BillingAccounts/1234/departments/42425",
                null,
                null,
                "usageStart eq '2018-04-01' AND usageEnd eq '2018-05-30'",
                null,
                Context.NONE);
    }
}
