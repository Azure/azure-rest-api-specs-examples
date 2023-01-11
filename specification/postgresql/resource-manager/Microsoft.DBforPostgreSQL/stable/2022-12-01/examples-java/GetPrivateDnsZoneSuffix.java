/** Samples for GetPrivateDnsZoneSuffix Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/examples/GetPrivateDnsZoneSuffix.json
     */
    /**
     * Sample code: GetPrivateDnsZoneSuffix.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getPrivateDnsZoneSuffix(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.getPrivateDnsZoneSuffixes().executeWithResponse(com.azure.core.util.Context.NONE);
    }
}
