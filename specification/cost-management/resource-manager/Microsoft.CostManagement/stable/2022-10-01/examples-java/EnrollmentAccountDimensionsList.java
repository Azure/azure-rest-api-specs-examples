/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/EnrollmentAccountDimensionsList.json
     */
    /**
     * Sample code: EnrollmentAccountDimensionsList-Legacy.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void enrollmentAccountDimensionsListLegacy(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
