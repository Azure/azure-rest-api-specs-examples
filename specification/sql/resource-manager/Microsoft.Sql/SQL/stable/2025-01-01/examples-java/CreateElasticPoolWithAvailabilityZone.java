
import com.azure.resourcemanager.sql.fluent.models.ElasticPoolInner;
import com.azure.resourcemanager.sql.models.AvailabilityZoneType;
import com.azure.resourcemanager.sql.models.ElasticPoolPerDatabaseSettings;
import com.azure.resourcemanager.sql.models.Sku;

/**
 * Samples for ElasticPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateElasticPoolWithAvailabilityZone.json
     */
    /**
     * Sample code: Create or Update an elastic pool with Availability Zone.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createOrUpdateAnElasticPoolWithAvailabilityZone(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().createOrUpdate("sqlcrudtest-2369", "sqlcrudtest-8069",
            "sqlcrudtest-8102",
            new ElasticPoolInner().withLocation("Japan East").withSku(new Sku().withName("HS_Gen5_4"))
                .withPerDatabaseSettings(
                    new ElasticPoolPerDatabaseSettings().withMinCapacity(0.25D).withMaxCapacity(2.0D))
                .withZoneRedundant(true).withAvailabilityZone(AvailabilityZoneType.ONE),
            com.azure.core.util.Context.NONE);
    }
}
