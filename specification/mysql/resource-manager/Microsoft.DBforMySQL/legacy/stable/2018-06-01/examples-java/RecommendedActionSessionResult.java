import com.azure.core.util.Context;

/** Samples for LocationBasedRecommendedActionSessionsResult List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/RecommendedActionSessionResult.json
     */
    /**
     * Sample code: RecommendedActionSessionResult.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void recommendedActionSessionResult(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .locationBasedRecommendedActionSessionsResults()
            .list("WestUS", "aaaabbbb-cccc-dddd-0000-111122223333", Context.NONE);
    }
}
