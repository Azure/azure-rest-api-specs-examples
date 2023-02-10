/** Samples for Exports List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportsGetByEnrollmentAccount.json
     */
    /**
     * Sample code: ExportsGetByEnrollmentAccount.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportsGetByEnrollmentAccount(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .listWithResponse(
                "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456",
                null,
                com.azure.core.util.Context.NONE);
    }
}
