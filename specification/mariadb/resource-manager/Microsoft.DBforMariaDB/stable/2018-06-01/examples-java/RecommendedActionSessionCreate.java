/** Samples for ResourceProvider CreateRecommendedActionSession. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/RecommendedActionSessionCreate.json
     */
    /**
     * Sample code: RecommendedActionSessionCreate.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void recommendedActionSessionCreate(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .resourceProviders()
            .createRecommendedActionSession(
                "testResourceGroupName",
                "testServerName",
                "Index",
                "someDatabaseName",
                com.azure.core.util.Context.NONE);
    }
}
