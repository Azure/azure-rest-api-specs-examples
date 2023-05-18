import com.azure.resourcemanager.mysqlflexibleserver.models.ServerGtidSetParameter;

/** Samples for Servers ResetGtid. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2022-09-30-preview/examples/ServerResetGtid.json
     */
    /**
     * Sample code: Reset GTID on a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void resetGTIDOnAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager
            .servers()
            .resetGtid(
                "TestGroup",
                "testserver",
                new ServerGtidSetParameter().withGtidSet("4aff5b51-97ba-11ed-a955-002248036acc:1-16"),
                com.azure.core.util.Context.NONE);
    }
}
