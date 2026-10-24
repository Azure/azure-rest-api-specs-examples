
import com.azure.resourcemanager.horizondb.models.LogCaptureConfigurationForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.LogCaptureConfigurationPropertiesForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.LogCaptureState;

/**
 * Samples for LogCaptureConfigurations Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/LogCaptureConfigurations_Update.json
     */
    /**
     * Sample code: Disable PostgreSQL server log capture.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void disablePostgreSQLServerLogCapture(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.logCaptureConfigurations().update("example-resource-group", "example-cluster", "PostgreSqlServer",
            new LogCaptureConfigurationForPatchUpdate().withProperties(
                new LogCaptureConfigurationPropertiesForPatchUpdate().withState(LogCaptureState.DISABLED)),
            com.azure.core.util.Context.NONE);
    }
}
