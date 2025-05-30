
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionEnum;

/**
 * Samples for TuningConfiguration ListSessions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * TuningConfiguration_ListSessions.json
     */
    /**
     * Sample code: TuningConfiguration_ListSessions.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        tuningConfigurationListSessions(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningConfigurations().listSessions("testrg", "testserver", TuningOptionEnum.CONFIGURATION,
            com.azure.core.util.Context.NONE);
    }
}
