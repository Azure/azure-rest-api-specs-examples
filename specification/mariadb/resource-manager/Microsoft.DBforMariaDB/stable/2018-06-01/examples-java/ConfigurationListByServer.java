/** Samples for Configurations ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ConfigurationListByServer.json
     */
    /**
     * Sample code: ConfigurationList.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void configurationList(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.configurations().listByServer("testrg", "mariadbtestsvc1", com.azure.core.util.Context.NONE);
    }
}
