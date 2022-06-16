import com.azure.core.util.Context;

/** Samples for QueryTexts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/QueryTextsGet.json
     */
    /**
     * Sample code: QueryTextsGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void queryTextsGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.queryTexts().getWithResponse("testResourceGroupName", "testServerName", "1", Context.NONE);
    }
}
