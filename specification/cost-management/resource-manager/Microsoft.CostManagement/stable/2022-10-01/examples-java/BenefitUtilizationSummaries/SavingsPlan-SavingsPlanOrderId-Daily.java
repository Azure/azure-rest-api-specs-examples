/** Samples for BenefitUtilizationSummaries ListBySavingsPlanOrder. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/BenefitUtilizationSummaries/SavingsPlan-SavingsPlanOrderId-Daily.json
     */
    /**
     * Sample code: SavingsPlanUtilizationSummariesDaily.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void savingsPlanUtilizationSummariesDaily(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .benefitUtilizationSummaries()
            .listBySavingsPlanOrder(
                "66cccc66-6ccc-6c66-666c-66cc6c6c66c6", null, null, com.azure.core.util.Context.NONE);
    }
}
