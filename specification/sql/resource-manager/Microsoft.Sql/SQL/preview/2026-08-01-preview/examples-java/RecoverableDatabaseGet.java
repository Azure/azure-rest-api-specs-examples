
/**
 * Samples for RecoverableDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/RecoverableDatabaseGet.json
     */
    /**
     * Sample code: Get a recoverable database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getARecoverableDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRecoverableDatabases().getWithResponse("recoverabledatabasetest-6852",
            "recoverabledatabasetest-2080", "recoverabledatabasetest-9187", null, null,
            com.azure.core.util.Context.NONE);
    }
}
