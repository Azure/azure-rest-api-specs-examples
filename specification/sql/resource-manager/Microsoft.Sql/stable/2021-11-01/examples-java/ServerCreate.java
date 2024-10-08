
import com.azure.resourcemanager.sql.fluent.models.ServerInner;
import com.azure.resourcemanager.sql.models.PrincipalType;
import com.azure.resourcemanager.sql.models.ServerExternalAdministrator;
import com.azure.resourcemanager.sql.models.ServerNetworkAccessFlag;
import java.util.UUID;

/**
 * Samples for Servers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerCreate.json
     */
    /**
     * Sample code: Create server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServers().createOrUpdate("sqlcrudtest-7398", "sqlcrudtest-4645",
            new ServerInner().withLocation("Japan East").withAdministratorLogin("dummylogin")
                .withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withPublicNetworkAccess(ServerNetworkAccessFlag.ENABLED)
                .withAdministrators(new ServerExternalAdministrator().withPrincipalType(PrincipalType.USER)
                    .withLogin("bob@contoso.com").withSid(UUID.fromString("00000011-1111-2222-2222-123456789111"))
                    .withTenantId(UUID.fromString("00000011-1111-2222-2222-123456789111"))
                    .withAzureADOnlyAuthentication(true))
                .withRestrictOutboundNetworkAccess(ServerNetworkAccessFlag.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
