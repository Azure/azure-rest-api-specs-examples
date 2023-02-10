/** Samples for Exports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportGetByResourceGroup.json
     */
    /**
     * Sample code: ExportGetByResourceGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportGetByResourceGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG",
                "TestExport",
                null,
                com.azure.core.util.Context.NONE);
    }
}
