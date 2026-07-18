
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;

/**
 * Samples for ElasticPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ElasticPoolUpdateAssignMaintenanceConfiguration.json
     */
    /**
     * Sample code: Assigns maintenance configuration to an elastic pool.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        assignsMaintenanceConfigurationToAnElasticPool(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().update("sqlcrudtest-2369", "sqlcrudtest-8069", "sqlcrudtest-8102",
            new ElasticPoolUpdate().withMaintenanceConfigurationId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_1"),
            com.azure.core.util.Context.NONE);
    }
}
