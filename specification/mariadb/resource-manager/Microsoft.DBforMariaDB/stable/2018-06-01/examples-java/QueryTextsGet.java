/** Samples for QueryTexts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/QueryTextsGet.json
     */
    /**
     * Sample code: QueryTextsGet.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void queryTextsGet(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .queryTexts()
            .getWithResponse("testResourceGroupName", "testServerName", "1", com.azure.core.util.Context.NONE);
    }
}
