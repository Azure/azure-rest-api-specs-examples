import com.azure.resourcemanager.costmanagement.models.CostDetailsMetricType;
import com.azure.resourcemanager.costmanagement.models.GenerateCostDetailsReportRequestDefinition;

/** Samples for GenerateCostDetailsReport CreateOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateCostDetailsReportByBillingProfileAndInvoiceId.json
     */
    /**
     * Sample code: GenerateCostDetailsReportByBillingProfileAndInvoiceId.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void generateCostDetailsReportByBillingProfileAndInvoiceId(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .generateCostDetailsReports()
            .createOperation(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579",
                new GenerateCostDetailsReportRequestDefinition()
                    .withMetric(CostDetailsMetricType.ACTUAL_COST)
                    .withInvoiceId("M1234567"),
                com.azure.core.util.Context.NONE);
    }
}
