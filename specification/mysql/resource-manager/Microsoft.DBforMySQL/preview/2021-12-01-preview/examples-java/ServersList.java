/** Samples for Servers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/preview/2021-12-01-preview/examples/ServersList.json
     */
    /**
     * Sample code: List servers in a subscription.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listServersInASubscription(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().list(com.azure.core.util.Context.NONE);
    }
}
