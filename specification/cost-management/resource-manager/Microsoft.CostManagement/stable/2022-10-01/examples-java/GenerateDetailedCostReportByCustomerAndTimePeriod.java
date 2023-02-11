import com.azure.resourcemanager.costmanagement.models.GenerateDetailedCostReportDefinition;
import com.azure.resourcemanager.costmanagement.models.GenerateDetailedCostReportMetricType;
import com.azure.resourcemanager.costmanagement.models.GenerateDetailedCostReportTimePeriod;

/** Samples for GenerateDetailedCostReport CreateOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateDetailedCostReportByCustomerAndTimePeriod.json
     */
    /**
     * Sample code: GenerateDetailedCostReportByCustomerAndTimePeriod.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void generateDetailedCostReportByCustomerAndTimePeriod(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .generateDetailedCostReports()
            .createOperation(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/customers/13579",
                new GenerateDetailedCostReportDefinition()
                    .withMetric(GenerateDetailedCostReportMetricType.ACTUAL_COST)
                    .withTimePeriod(
                        new GenerateDetailedCostReportTimePeriod().withStart("2020-03-01").withEnd("2020-03-15")),
                com.azure.core.util.Context.NONE);
    }
}
