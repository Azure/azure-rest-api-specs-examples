/** Samples for BenefitRecommendations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/BenefitRecommendationsByBillingAccount.json
     */
    /**
     * Sample code: BenefitRecommendationsBillingAccountList.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void benefitRecommendationsBillingAccountList(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .benefitRecommendations()
            .list(
                "providers/Microsoft.Billing/billingAccounts/123456",
                "properties/lookBackPeriod eq 'Last7Days' AND properties/term eq 'P1Y'",
                null,
                "properties/usage,properties/allRecommendationDetails",
                com.azure.core.util.Context.NONE);
    }
}
