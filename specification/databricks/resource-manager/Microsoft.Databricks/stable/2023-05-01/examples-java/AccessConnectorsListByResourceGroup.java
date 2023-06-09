/** Samples for AccessConnectors ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-05-01/examples/AccessConnectorsListByResourceGroup.json
     */
    /**
     * Sample code: Lists azure databricks accessConnectors within a resource group.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void listsAzureDatabricksAccessConnectorsWithinAResourceGroup(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.accessConnectors().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
