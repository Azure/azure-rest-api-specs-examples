/** Samples for Exports List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportsGetByResourceGroup.json
     */
    /**
     * Sample code: ExportsGetByResourceGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportsGetByResourceGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .listWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG",
                null,
                com.azure.core.util.Context.NONE);
    }
}
