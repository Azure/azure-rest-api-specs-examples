import com.azure.resourcemanager.costmanagement.models.GenerateDetailedCostReportDefinition;
import com.azure.resourcemanager.costmanagement.models.GenerateDetailedCostReportMetricType;
import com.azure.resourcemanager.costmanagement.models.GenerateDetailedCostReportTimePeriod;

/** Samples for GenerateDetailedCostReport CreateOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateDetailedCostReportBySubscriptionAndTimePeriod.json
     */
    /**
     * Sample code: GenerateDetailedCostReportBySubscriptionAndTimePeriod.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void generateDetailedCostReportBySubscriptionAndTimePeriod(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .generateDetailedCostReports()
            .createOperation(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                new GenerateDetailedCostReportDefinition()
                    .withMetric(GenerateDetailedCostReportMetricType.ACTUAL_COST)
                    .withTimePeriod(
                        new GenerateDetailedCostReportTimePeriod().withStart("2020-03-01").withEnd("2020-03-15")),
                com.azure.core.util.Context.NONE);
    }
}
