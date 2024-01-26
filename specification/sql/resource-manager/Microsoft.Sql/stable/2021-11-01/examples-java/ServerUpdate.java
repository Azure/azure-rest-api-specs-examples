
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ServerNetworkAccessFlag;
import com.azure.resourcemanager.sql.models.ServerUpdate;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerUpdate.json
     */
    /**
     * Sample code: Update a server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServers().update("sqlcrudtest-7398", "sqlcrudtest-4645",
            new ServerUpdate().withAdministratorLogin("dummylogin")
                .withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withPublicNetworkAccess(ServerNetworkAccessFlag.DISABLED)
                .withRestrictOutboundNetworkAccess(ServerNetworkAccessFlag.ENABLED),
            Context.NONE);
    }
}
