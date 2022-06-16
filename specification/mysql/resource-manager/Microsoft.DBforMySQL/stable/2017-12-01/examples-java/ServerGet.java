import com.azure.core.util.Context;

/** Samples for Servers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerGet.json
     */
    /**
     * Sample code: ServerGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "mysqltestsvc4", Context.NONE);
    }
}
