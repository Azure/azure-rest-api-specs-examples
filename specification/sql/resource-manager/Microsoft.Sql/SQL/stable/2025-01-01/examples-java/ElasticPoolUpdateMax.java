
import com.azure.resourcemanager.sql.models.ElasticPoolLicenseType;
import com.azure.resourcemanager.sql.models.ElasticPoolPerDatabaseSettings;
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;
import com.azure.resourcemanager.sql.models.Sku;

/**
 * Samples for ElasticPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ElasticPoolUpdateMax.json
     */
    /**
     * Sample code: Update an elastic pool with all parameter.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAnElasticPoolWithAllParameter(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().update("sqlcrudtest-2369", "sqlcrudtest-8069", "sqlcrudtest-8102",
            new ElasticPoolUpdate().withSku(new Sku().withName("BC_Gen4").withTier("BusinessCritical").withCapacity(2))
                .withPerDatabaseSettings(
                    new ElasticPoolPerDatabaseSettings().withMinCapacity(0.25D).withMaxCapacity(1.0D))
                .withZoneRedundant(true).withLicenseType(ElasticPoolLicenseType.LICENSE_INCLUDED),
            com.azure.core.util.Context.NONE);
    }
}
