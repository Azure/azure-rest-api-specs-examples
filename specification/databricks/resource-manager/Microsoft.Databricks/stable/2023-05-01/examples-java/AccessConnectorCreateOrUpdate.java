/** Samples for AccessConnectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-05-01/examples/AccessConnectorCreateOrUpdate.json
     */
    /**
     * Sample code: Create an azure databricks accessConnector with SystemAssigned Identity.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void createAnAzureDatabricksAccessConnectorWithSystemAssignedIdentity(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .accessConnectors()
            .define("myAccessConnector")
            .withRegion("westus")
            .withExistingResourceGroup("rg")
            .create();
    }
}
