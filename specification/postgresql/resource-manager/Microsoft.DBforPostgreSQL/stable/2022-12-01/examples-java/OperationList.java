/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/examples/OperationList.json
     */
    /**
     * Sample code: OperationList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void operationList(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
