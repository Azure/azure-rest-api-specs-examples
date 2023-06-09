/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/OperationsList.json
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
