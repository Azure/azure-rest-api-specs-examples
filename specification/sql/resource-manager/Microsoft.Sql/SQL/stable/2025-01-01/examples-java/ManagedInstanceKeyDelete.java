
/**
 * Samples for ManagedInstanceKeys Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceKeyDelete.json
     */
    /**
     * Sample code: Delete the managed instance key.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteTheManagedInstanceKey(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceKeys().delete("sqlcrudtest-7398", "sqlcrudtest-4645",
            "someVault_someKey_01234567890123456789012345678901", com.azure.core.util.Context.NONE);
    }
}
