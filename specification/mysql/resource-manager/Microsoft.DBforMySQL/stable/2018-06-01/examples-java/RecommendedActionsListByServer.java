import com.azure.core.util.Context;

/** Samples for RecommendedActions ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/RecommendedActionsListByServer.json
     */
    /**
     * Sample code: RecommendedActionsListByServer.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void recommendedActionsListByServer(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .recommendedActions()
            .listByServer("testResourceGroupName", "testServerName", "Index", null, Context.NONE);
    }
}
