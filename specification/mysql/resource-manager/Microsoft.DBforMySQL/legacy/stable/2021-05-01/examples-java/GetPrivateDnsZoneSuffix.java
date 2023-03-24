import com.azure.core.util.Context;

/** Samples for GetPrivateDnsZoneSuffix Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/GetPrivateDnsZoneSuffix.json
     */
    /**
     * Sample code: GetPrivateDnsZoneSuffix.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getPrivateDnsZoneSuffix(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.getPrivateDnsZoneSuffixes().executeWithResponse(Context.NONE);
    }
}
