/** Samples for BenefitUtilizationSummaries ListBySavingsPlanId. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/BenefitUtilizationSummaries/SavingsPlan-SavingsPlanId-Monthly.json
     */
    /**
     * Sample code: SavingsPlanUtilizationSummariesMonthlyWithSavingsPlanId.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void savingsPlanUtilizationSummariesMonthlyWithSavingsPlanId(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .benefitUtilizationSummaries()
            .listBySavingsPlanId(
                "66cccc66-6ccc-6c66-666c-66cc6c6c66c6",
                "222d22dd-d2d2-2dd2-222d-2dd2222ddddd",
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
