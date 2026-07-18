
import com.azure.resourcemanager.sql.fluent.models.ElasticPoolInner;

/**
 * Samples for ElasticPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ElasticPoolCreateOrUpdateSetMaintenanceConfiguration.json
     */
    /**
     * Sample code: Create or update elastic pool with maintenance configuration parameter.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateElasticPoolWithMaintenanceConfigurationParameter(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().createOrUpdate("sqlcrudtest-2369", "sqlcrudtest-8069",
            "sqlcrudtest-8102",
            new ElasticPoolInner().withLocation("Japan East").withMaintenanceConfigurationId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_1"),
            com.azure.core.util.Context.NONE);
    }
}
