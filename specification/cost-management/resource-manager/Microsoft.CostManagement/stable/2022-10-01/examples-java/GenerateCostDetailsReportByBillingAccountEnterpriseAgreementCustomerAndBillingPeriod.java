import com.azure.resourcemanager.costmanagement.models.CostDetailsMetricType;
import com.azure.resourcemanager.costmanagement.models.GenerateCostDetailsReportRequestDefinition;

/** Samples for GenerateCostDetailsReport CreateOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateCostDetailsReportByBillingAccountEnterpriseAgreementCustomerAndBillingPeriod.json
     */
    /**
     * Sample code: GenerateCostDetailsReportByBillingAccountEnterpriseAgreementCustomerAndBillingPeriod.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void generateCostDetailsReportByBillingAccountEnterpriseAgreementCustomerAndBillingPeriod(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .generateCostDetailsReports()
            .createOperation(
                "providers/Microsoft.Billing/billingAccounts/12345",
                new GenerateCostDetailsReportRequestDefinition()
                    .withMetric(CostDetailsMetricType.ACTUAL_COST)
                    .withBillingPeriod("202205"),
                com.azure.core.util.Context.NONE);
    }
}
