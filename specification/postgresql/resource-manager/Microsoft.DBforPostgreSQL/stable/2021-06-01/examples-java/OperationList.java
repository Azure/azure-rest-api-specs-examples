import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/OperationList.json
     */
    /**
     * Sample code: OperationList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void operationList(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.operations().listWithResponse(Context.NONE);
    }
}
