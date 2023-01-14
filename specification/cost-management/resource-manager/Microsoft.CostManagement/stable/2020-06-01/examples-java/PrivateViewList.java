/** Samples for Views List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/PrivateViewList.json
     */
    /**
     * Sample code: PrivateViewList.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void privateViewList(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.views().list(com.azure.core.util.Context.NONE);
    }
}
