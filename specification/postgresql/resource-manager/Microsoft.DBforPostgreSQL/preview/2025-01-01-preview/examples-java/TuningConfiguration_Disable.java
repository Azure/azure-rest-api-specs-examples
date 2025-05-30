
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionEnum;

/**
 * Samples for TuningConfiguration Disable.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * TuningConfiguration_Disable.json
     */
    /**
     * Sample code: TuningConfiguration_Disable.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        tuningConfigurationDisable(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningConfigurations().disable("testrg", "testserver", TuningOptionEnum.CONFIGURATION,
            com.azure.core.util.Context.NONE);
    }
}
