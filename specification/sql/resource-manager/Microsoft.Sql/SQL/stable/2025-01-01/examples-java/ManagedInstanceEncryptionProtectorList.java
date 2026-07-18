
/**
 * Samples for ManagedInstanceEncryptionProtectors ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceEncryptionProtectorList.json
     */
    /**
     * Sample code: List encryption protectors by managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listEncryptionProtectorsByManagedInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceEncryptionProtectors().listByInstance("sqlcrudtest-7398",
            "sqlcrudtest-4645", com.azure.core.util.Context.NONE);
    }
}
