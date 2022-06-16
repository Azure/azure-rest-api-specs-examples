import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/OperationsList.json
     */
    /**
     * Sample code: OperationList.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void operationList(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.operations().list(Context.NONE);
    }
}
