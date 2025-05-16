
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionEnum;

/**
 * Samples for TuningConfiguration ListSessionDetails.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * TuningConfiguration_ListSessionDetails.json
     */
    /**
     * Sample code: TuningConfiguration_ListSessionDetails.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void tuningConfigurationListSessionDetails(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningConfigurations().listSessionDetails("testrg", "testserver", TuningOptionEnum.CONFIGURATION,
            "oooooooo-oooo-oooo-oooo-oooooooooooo", com.azure.core.util.Context.NONE);
    }
}
