
import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Gets a private link resource for SQL.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAPrivateLinkResourceForSQL(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getPrivateLinkResources().getWithResponse("Default", "test-svr",
            "plr", Context.NONE);
    }
}
