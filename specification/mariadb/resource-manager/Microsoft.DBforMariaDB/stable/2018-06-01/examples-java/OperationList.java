/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/OperationList.json
     */
    /**
     * Sample code: OperationList.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void operationList(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
