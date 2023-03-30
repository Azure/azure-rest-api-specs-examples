import com.azure.core.util.Context;

/** Samples for LocationBasedRecommendedActionSessionsOperationStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/RecommendedActionSessionOperationStatus.json
     */
    /**
     * Sample code: RecommendedActionSessionOperationStatus.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void recommendedActionSessionOperationStatus(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .locationBasedRecommendedActionSessionsOperationStatus()
            .getWithResponse("WestUS", "aaaabbbb-cccc-dddd-0000-111122223333", Context.NONE);
    }
}
