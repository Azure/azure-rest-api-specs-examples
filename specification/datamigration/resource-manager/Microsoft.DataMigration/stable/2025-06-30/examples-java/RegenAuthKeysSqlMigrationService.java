
import com.azure.resourcemanager.datamigration.fluent.models.RegenAuthKeysInner;

/**
 * Samples for SqlMigrationServices RegenerateAuthKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * RegenAuthKeysSqlMigrationService.json
     */
    /**
     * Sample code: Regenerate the of Authentication Keys.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void
        regenerateTheOfAuthenticationKeys(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.sqlMigrationServices().regenerateAuthKeysWithResponse("testrg", "service1",
            new RegenAuthKeysInner().withKeyName("fakeTokenPlaceholder"), com.azure.core.util.Context.NONE);
    }
}
