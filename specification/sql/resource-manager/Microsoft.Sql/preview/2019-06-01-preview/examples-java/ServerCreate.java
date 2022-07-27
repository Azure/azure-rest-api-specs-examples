import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ServerInner;

/** Samples for Servers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/ServerCreate.json
     */
    /**
     * Sample code: Create server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServers()
            .createOrUpdate(
                "sqlcrudtest-7398",
                "sqlcrudtest-4645",
                new ServerInner()
                    .withLocation("Japan East")
                    .withAdministratorLogin("dummylogin")
                    .withAdministratorLoginPassword("Un53cuRE!"),
                Context.NONE);
    }
}
