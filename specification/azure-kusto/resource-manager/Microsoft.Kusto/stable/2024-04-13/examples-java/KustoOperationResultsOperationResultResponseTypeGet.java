
/**
 * Samples for OperationsResultsLocation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/
     * KustoOperationResultsOperationResultResponseTypeGet.json
     */
    /**
     * Sample code: KustoOperationsResultsLocationGet.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoOperationsResultsLocationGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.operationsResultsLocations().getWithResponse("westus", "30972f1b-b61d-4fd8-bd34-3dcfa24670f3",
            com.azure.core.util.Context.NONE);
    }
}
