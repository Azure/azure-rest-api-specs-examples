
import java.util.Arrays;

/**
 * Samples for QueryTexts ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/QueryTextsListByServer.
     * json
     */
    /**
     * Sample code: QueryTextsListByServer.
     * 
     * @param manager Entry point to MariaDBManager.
     */
    public static void queryTextsListByServer(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.queryTexts().listByServer("testResourceGroupName", "testServerName", Arrays.asList("1", "2"),
            com.azure.core.util.Context.NONE);
    }
}
