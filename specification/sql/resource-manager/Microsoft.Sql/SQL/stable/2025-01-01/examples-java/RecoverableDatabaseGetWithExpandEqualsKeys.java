
/**
 * Samples for RecoverableDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/RecoverableDatabaseGetWithExpandEqualsKeys.json
     */
    /**
     * Sample code: Gets a recoverable database with expand equals keys.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsARecoverableDatabaseWithExpandEqualsKeys(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRecoverableDatabases().getWithResponse("recoverabledatabasetest-6852",
            "recoverabledatabasetest-2080", "recoverabledatabasetest-9187", "keys", null,
            com.azure.core.util.Context.NONE);
    }
}
