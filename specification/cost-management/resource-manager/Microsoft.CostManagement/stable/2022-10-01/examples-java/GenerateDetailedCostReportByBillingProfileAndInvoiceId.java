import com.azure.resourcemanager.costmanagement.models.GenerateDetailedCostReportDefinition;
import com.azure.resourcemanager.costmanagement.models.GenerateDetailedCostReportMetricType;

/** Samples for GenerateDetailedCostReport CreateOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateDetailedCostReportByBillingProfileAndInvoiceId.json
     */
    /**
     * Sample code: GenerateDetailedCostReportByBillingProfileAndInvoiceId.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void generateDetailedCostReportByBillingProfileAndInvoiceId(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .generateDetailedCostReports()
            .createOperation(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579",
                new GenerateDetailedCostReportDefinition()
                    .withMetric(GenerateDetailedCostReportMetricType.ACTUAL_COST)
                    .withInvoiceId("M1234567"),
                com.azure.core.util.Context.NONE);
    }
}
