/** Samples for AccessConnectors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-05-01/examples/AccessConnectorsListBySubscriptionId.json
     */
    /**
     * Sample code: Lists all the azure databricks accessConnectors within a subscription.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void listsAllTheAzureDatabricksAccessConnectorsWithinASubscription(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.accessConnectors().list(com.azure.core.util.Context.NONE);
    }
}
