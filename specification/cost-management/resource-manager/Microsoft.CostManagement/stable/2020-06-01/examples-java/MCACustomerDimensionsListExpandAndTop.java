/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/MCACustomerDimensionsListExpandAndTop.json
     */
    /**
     * Sample code: CustomerDimensionsListExpandAndTop-Modern.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void customerDimensionsListExpandAndTopModern(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/customers/5678",
                null,
                "properties/data",
                null,
                5,
                com.azure.core.util.Context.NONE);
    }
}
