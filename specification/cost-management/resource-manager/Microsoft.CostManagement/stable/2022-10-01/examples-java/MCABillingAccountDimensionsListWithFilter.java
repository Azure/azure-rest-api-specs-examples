
/**
 * Samples for Dimensions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/
     * MCABillingAccountDimensionsListWithFilter.json
     */
    /**
     * Sample code: BillingAccountDimensionsListWithFilter-MCA.
     * 
     * @param manager Entry point to CostManagementManager.
     */
    public static void billingAccountDimensionsListWithFilterMCA(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.dimensions().list("providers/Microsoft.Billing/billingAccounts/12345:6789",
            "properties/category eq 'resourceId'", "properties/data", null, 5, com.azure.core.util.Context.NONE);
    }
}
