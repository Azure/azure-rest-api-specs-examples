
import com.azure.resourcemanager.redisenterprise.fluent.models.MigrationInner;
import com.azure.resourcemanager.redisenterprise.models.AzureCacheForRedisMigrationProperties;

/**
 * Samples for Migrations Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/RedisEnterpriseMigrationStart.json
     */
    /**
     * Sample code: RedisEnterpriseMigrationStart.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void
        redisEnterpriseMigrationStart(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.migrations().start("rg1", "cache1",
            new MigrationInner().withProperties(new AzureCacheForRedisMigrationProperties().withSourceResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1")
                .withSwitchDns(true).withSkipDataMigration(true)),
            com.azure.core.util.Context.NONE);
    }
}
