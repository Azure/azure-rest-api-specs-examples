
/**
 * Samples for ManagedInstanceKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceKeyGet.json
     */
    /**
     * Sample code: Get the managed instance key.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getTheManagedInstanceKey(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceKeys().getWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            "someVault_someKey_01234567890123456789012345678901", com.azure.core.util.Context.NONE);
    }
}
