/** Samples for Views Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/PrivateView.json
     */
    /**
     * Sample code: PrivateView.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void privateView(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.views().getWithResponse("swaggerExample", com.azure.core.util.Context.NONE);
    }
}
