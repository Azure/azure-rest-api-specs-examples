
/**
 * Samples for RecoverableDatabases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/RecoverableDatabaseList.json
     */
    /**
     * Sample code: Get list of recoverable databases.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getListOfRecoverableDatabases(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRecoverableDatabases().listByServer("recoverabledatabasetest-1234",
            "recoverabledatabasetest-7177", com.azure.core.util.Context.NONE);
    }
}
