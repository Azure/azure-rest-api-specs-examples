
/**
 * Samples for Dimensions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/
     * MCAInvoiceSectionDimensionsListWithFilter.json
     */
    /**
     * Sample code: InvoiceSectionDimensionsListWithFilter-MCA.
     * 
     * @param manager Entry point to CostManagementManager.
     */
    public static void invoiceSectionDimensionsListWithFilterMCA(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.dimensions().list(
            "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/invoiceSections/9876",
            "properties/category eq 'resourceId'", "properties/data", null, 5, com.azure.core.util.Context.NONE);
    }
}
