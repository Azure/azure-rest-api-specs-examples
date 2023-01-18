/** Samples for RecommendedActions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/RecommendedActionsGet.json
     */
    /**
     * Sample code: RecommendedActionsGet.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void recommendedActionsGet(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .recommendedActions()
            .getWithResponse(
                "testResourceGroupName", "testServerName", "Index", "Index-1", com.azure.core.util.Context.NONE);
    }
}
