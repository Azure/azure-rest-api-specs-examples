/** Samples for Advisors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/AdvisorsGet.json
     */
    /**
     * Sample code: AdvisorsGet.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void advisorsGet(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .advisors()
            .getWithResponse("testResourceGroupName", "testServerName", "Index", com.azure.core.util.Context.NONE);
    }
}
