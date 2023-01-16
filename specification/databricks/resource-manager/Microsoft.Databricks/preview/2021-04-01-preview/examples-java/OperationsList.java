/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/OperationsList.json
     */
    /**
     * Sample code: Operations.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void operations(com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
