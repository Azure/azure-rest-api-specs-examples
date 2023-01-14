/** Samples for Views Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/PrivateViewDelete.json
     */
    /**
     * Sample code: DeletePrivateView.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void deletePrivateView(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.views().deleteWithResponse("TestView", com.azure.core.util.Context.NONE);
    }
}
