/** Samples for Servers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerGet.json
     */
    /**
     * Sample code: ServerGet.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void serverGet(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "mariadbtestsvc4", com.azure.core.util.Context.NONE);
    }
}
