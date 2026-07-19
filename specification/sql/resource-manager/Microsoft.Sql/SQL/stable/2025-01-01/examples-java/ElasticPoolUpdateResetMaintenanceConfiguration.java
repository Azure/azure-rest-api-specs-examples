
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;

/**
 * Samples for ElasticPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ElasticPoolUpdateResetMaintenanceConfiguration.json
     */
    /**
     * Sample code: Resets maintenance configuration of an elastic pool to default.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        resetsMaintenanceConfigurationOfAnElasticPoolToDefault(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().update("sqlcrudtest-2369", "sqlcrudtest-8069", "sqlcrudtest-8102",
            new ElasticPoolUpdate().withMaintenanceConfigurationId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_Default"),
            com.azure.core.util.Context.NONE);
    }
}
