
/**
 * Samples for Dimensions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/
     * MCACustomerDimensionsListWithFilter.json
     */
    /**
     * Sample code: CustomerDimensionsListWithFilter-MCA.
     * 
     * @param manager Entry point to CostManagementManager.
     */
    public static void
        customerDimensionsListWithFilterMCA(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.dimensions().list("providers/Microsoft.Billing/billingAccounts/12345:6789/customers/5678",
            "properties/category eq 'resourceId'", "properties/data", null, 5, com.azure.core.util.Context.NONE);
    }
}
