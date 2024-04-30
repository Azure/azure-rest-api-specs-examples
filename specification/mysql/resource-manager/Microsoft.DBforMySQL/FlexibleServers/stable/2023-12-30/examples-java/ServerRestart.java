
import com.azure.resourcemanager.mysqlflexibleserver.models.EnableStatusEnum;
import com.azure.resourcemanager.mysqlflexibleserver.models.ServerRestartParameter;

/**
 * Samples for Servers Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * ServerRestart.json
     */
    /**
     * Sample code: Restart a server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void restartAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().restart("TestGroup", "testserver",
            new ServerRestartParameter().withRestartWithFailover(EnableStatusEnum.ENABLED).withMaxFailoverSeconds(60),
            com.azure.core.util.Context.NONE);
    }
}
