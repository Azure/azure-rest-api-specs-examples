
/**
 * Samples for Administrators Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/AdministratorGet.
     * json
     */
    /**
     * Sample code: ServerGet.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverGet(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.administrators().getWithResponse("testrg", "pgtestsvc1", "oooooooo-oooo-oooo-oooo-oooooooooooo",
            com.azure.core.util.Context.NONE);
    }
}
