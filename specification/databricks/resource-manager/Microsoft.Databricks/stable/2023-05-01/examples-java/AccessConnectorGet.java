/** Samples for AccessConnectors GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-05-01/examples/AccessConnectorGet.json
     */
    /**
     * Sample code: Get an azure databricks accessConnector.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void getAnAzureDatabricksAccessConnector(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .accessConnectors()
            .getByResourceGroupWithResponse("rg", "myAccessConnector", com.azure.core.util.Context.NONE);
    }
}
