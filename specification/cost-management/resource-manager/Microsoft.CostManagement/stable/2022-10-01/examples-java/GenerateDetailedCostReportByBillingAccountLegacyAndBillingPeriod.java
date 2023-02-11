import com.azure.resourcemanager.costmanagement.models.GenerateDetailedCostReportDefinition;
import com.azure.resourcemanager.costmanagement.models.GenerateDetailedCostReportMetricType;

/** Samples for GenerateDetailedCostReport CreateOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateDetailedCostReportByBillingAccountLegacyAndBillingPeriod.json
     */
    /**
     * Sample code: GenerateDetailedCostReportByBillingAccountLegacyAndBillingPeriod.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void generateDetailedCostReportByBillingAccountLegacyAndBillingPeriod(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .generateDetailedCostReports()
            .createOperation(
                "providers/Microsoft.Billing/billingAccounts/12345",
                new GenerateDetailedCostReportDefinition()
                    .withMetric(GenerateDetailedCostReportMetricType.ACTUAL_COST)
                    .withBillingPeriod("202008"),
                com.azure.core.util.Context.NONE);
    }
}
