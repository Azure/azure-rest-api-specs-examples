
/**
 * Samples for Views GetByScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/
     * ViewByResourceGroup.json
     */
    /**
     * Sample code: ResourceGroupView.
     * 
     * @param manager Entry point to CostManagementManager.
     */
    public static void resourceGroupView(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.views().getByScopeWithResponse(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
