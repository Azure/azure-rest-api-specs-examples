
/**
 * Samples for OperationsResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoOperationResultsGet.
     * json
     */
    /**
     * Sample code: KustoOperationResultsGet.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoOperationResultsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.operationsResults().getWithResponse("westus", "30972f1b-b61d-4fd8-bd34-3dcfa24670f3",
            com.azure.core.util.Context.NONE);
    }
}
