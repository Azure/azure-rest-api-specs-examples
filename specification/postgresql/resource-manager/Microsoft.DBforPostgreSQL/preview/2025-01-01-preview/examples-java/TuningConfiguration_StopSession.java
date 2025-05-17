
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionEnum;

/**
 * Samples for TuningConfiguration StopSession.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * TuningConfiguration_StopSession.json
     */
    /**
     * Sample code: TuningConfiguration_StopSession.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        tuningConfigurationStopSession(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningConfigurations().stopSession("testrg", "testserver", TuningOptionEnum.CONFIGURATION,
            com.azure.core.util.Context.NONE);
    }
}
