/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param manager Entry point to OperationsManagementManager.
     */
    public static void operationsList(
        com.azure.resourcemanager.operationsmanagement.OperationsManagementManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
