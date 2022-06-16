import com.azure.core.util.Context;

/** Samples for SqlManagedInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/ListSubscriptionSqlManagedInstance.json
     */
    /**
     * Sample code: Gets all SQL Instance in a subscription.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void getsAllSQLInstanceInASubscription(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.sqlManagedInstances().list(Context.NONE);
    }
}
