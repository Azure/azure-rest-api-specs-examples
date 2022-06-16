import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/OperationList.json
     */
    /**
     * Sample code: OperationList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void operationList(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.operations().listWithResponse(Context.NONE);
    }
}
