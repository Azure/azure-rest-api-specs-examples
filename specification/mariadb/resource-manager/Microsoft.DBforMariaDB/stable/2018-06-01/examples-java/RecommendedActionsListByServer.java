/** Samples for RecommendedActions ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/RecommendedActionsListByServer.json
     */
    /**
     * Sample code: RecommendedActionsListByServer.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void recommendedActionsListByServer(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .recommendedActions()
            .listByServer("testResourceGroupName", "testServerName", "Index", null, com.azure.core.util.Context.NONE);
    }
}
