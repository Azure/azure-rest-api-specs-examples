import com.azure.core.util.Context;

/** Samples for RecommendedActions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/RecommendedActionsGet.json
     */
    /**
     * Sample code: RecommendedActionsGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void recommendedActionsGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .recommendedActions()
            .getWithResponse("testResourceGroupName", "testServerName", "Index", "Index-1", Context.NONE);
    }
}
