import com.azure.resourcemanager.costmanagement.models.CostDetailsMetricType;
import com.azure.resourcemanager.costmanagement.models.CostDetailsTimePeriod;
import com.azure.resourcemanager.costmanagement.models.GenerateCostDetailsReportRequestDefinition;

/** Samples for GenerateCostDetailsReport CreateOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateCostDetailsReportByCustomerAndTimePeriod.json
     */
    /**
     * Sample code: GenerateCostDetailsReportByCustomerAndTimePeriod.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void generateCostDetailsReportByCustomerAndTimePeriod(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .generateCostDetailsReports()
            .createOperation(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/customers/13579",
                new GenerateCostDetailsReportRequestDefinition()
                    .withMetric(CostDetailsMetricType.ACTUAL_COST)
                    .withTimePeriod(new CostDetailsTimePeriod().withStart("2020-03-01").withEnd("2020-03-15")),
                com.azure.core.util.Context.NONE);
    }
}
