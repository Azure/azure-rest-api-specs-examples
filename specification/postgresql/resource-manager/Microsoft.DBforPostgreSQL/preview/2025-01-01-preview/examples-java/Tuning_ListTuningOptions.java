
/**
 * Samples for TuningOptions ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * Tuning_ListTuningOptions.json
     */
    /**
     * Sample code: TuningOptions_ListByServer.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        tuningOptionsListByServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningOptions().listByServer("testrg", "testserver", com.azure.core.util.Context.NONE);
    }
}
