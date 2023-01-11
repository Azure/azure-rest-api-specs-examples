/** Samples for SqlServerInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/ListSubscriptionSqlServerInstance.json
     */
    /**
     * Sample code: Gets all SQL Server Instance in a subscription.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void getsAllSQLServerInstanceInASubscription(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.sqlServerInstances().list(com.azure.core.util.Context.NONE);
    }
}
