/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/MCAInvoiceSectionDimensionsListExpandAndTop.json
     */
    /**
     * Sample code: InvoiceSectionDimensionsListExpandAndTop-Modern.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void invoiceSectionDimensionsListExpandAndTopModern(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/invoiceSections/9876",
                null,
                "properties/data",
                null,
                5,
                com.azure.core.util.Context.NONE);
    }
}
