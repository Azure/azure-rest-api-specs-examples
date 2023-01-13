/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/MCACustomerDimensionsList.json
     */
    /**
     * Sample code: CustomerDimensionsList-Modern.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void customerDimensionsListModern(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/customers/5678",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
