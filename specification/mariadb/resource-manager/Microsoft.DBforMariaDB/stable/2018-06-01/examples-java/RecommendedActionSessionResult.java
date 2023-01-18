/** Samples for LocationBasedRecommendedActionSessionsResult List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/RecommendedActionSessionResult.json
     */
    /**
     * Sample code: RecommendedActionSessionResult.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void recommendedActionSessionResult(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .locationBasedRecommendedActionSessionsResults()
            .list("WestUS", "aaaabbbb-cccc-dddd-0000-111122223333", com.azure.core.util.Context.NONE);
    }
}
