
import com.azure.resourcemanager.postgresqlflexibleserver.models.ConfigTuningRequestParameter;
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionEnum;

/**
 * Samples for TuningConfiguration StartSession.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * TuningConfiguration_StartSession.json
     */
    /**
     * Sample code: TuningConfiguration_StartSession.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        tuningConfigurationStartSession(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningConfigurations().startSession(
            "testrg", "testserver", TuningOptionEnum.CONFIGURATION, new ConfigTuningRequestParameter()
                .withAllowServerRestarts(false).withTargetImprovementMetric("targetImprovementMetric"),
            com.azure.core.util.Context.NONE);
    }
}
