import com.azure.core.util.Context;

/** Samples for Servers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/ServerGetWithVnet.json
     */
    /**
     * Sample code: ServerGetWithVnet.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverGetWithVnet(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "pgtestsvc4", Context.NONE);
    }
}
