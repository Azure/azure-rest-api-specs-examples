/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/MCAInvoiceSectionDimensionsList.json
     */
    /**
     * Sample code: InvoiceSectionDimensionsList-Modern.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void invoiceSectionDimensionsListModern(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/invoiceSections/9876",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
