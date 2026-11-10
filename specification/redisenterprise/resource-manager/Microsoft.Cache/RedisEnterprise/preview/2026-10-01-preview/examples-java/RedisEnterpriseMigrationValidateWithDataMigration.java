
import com.azure.resourcemanager.redisenterprise.models.MigrationValidationRequest;

/**
 * Samples for Migrations Validate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/RedisEnterpriseMigrationValidateWithDataMigration.json
     */
    /**
     * Sample code: RedisEnterpriseMigrationValidateWithDataMigration.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseMigrationValidateWithDataMigration(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.migrations().validateWithResponse("rg1", "cache1",
            new MigrationValidationRequest().withSourceResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1")
                .withSkipDataMigration(false),
            com.azure.core.util.Context.NONE);
    }
}
