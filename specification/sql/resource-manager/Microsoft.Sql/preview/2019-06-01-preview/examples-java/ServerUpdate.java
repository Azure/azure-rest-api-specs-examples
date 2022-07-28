import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ServerUpdate;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/ServerUpdate.json
     */
    /**
     * Sample code: Update a server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServers()
            .update(
                "sqlcrudtest-7398",
                "sqlcrudtest-4645",
                new ServerUpdate().withAdministratorLogin("dummylogin").withAdministratorLoginPassword("Un53cuRE!"),
                Context.NONE);
    }
}
