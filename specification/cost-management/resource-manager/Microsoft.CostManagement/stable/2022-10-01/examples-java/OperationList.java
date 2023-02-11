/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/OperationList.json
     */
    /**
     * Sample code: OperationList.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void operationList(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
