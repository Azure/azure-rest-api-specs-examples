
import com.azure.resourcemanager.redisenterprise.models.MigrationRollbackReason;
import com.azure.resourcemanager.redisenterprise.models.MigrationUndoRequest;

/**
 * Samples for Migrations Undo.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/RedisEnterpriseMigrationUndo.json
     */
    /**
     * Sample code: RedisEnterpriseMigrationUndo.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void
        redisEnterpriseMigrationUndo(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.migrations().undo("rg1", "cache1",
            new MigrationUndoRequest().withRollbackReason(MigrationRollbackReason.PERFORMANCE),
            com.azure.core.util.Context.NONE);
    }
}
