
import com.azure.resourcemanager.sql.fluent.models.ServerTrustGroupInner;
import com.azure.resourcemanager.sql.models.ServerInfo;
import com.azure.resourcemanager.sql.models.ServerTrustGroupPropertiesTrustScopesItem;
import java.util.Arrays;

/**
 * Samples for ServerTrustGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerTrustGroupCreate.json
     */
    /**
     * Sample code: Create server trust group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createServerTrustGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerTrustGroups().createOrUpdate("Default", "Japan East",
            "server-trust-group-test",
            new ServerTrustGroupInner().withGroupMembers(Arrays.asList(new ServerInfo().withServerId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/managedInstance-1"),
                new ServerInfo().withServerId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/managedInstance-2")))
                .withTrustScopes(Arrays.asList(ServerTrustGroupPropertiesTrustScopesItem.GLOBAL_TRANSACTIONS,
                    ServerTrustGroupPropertiesTrustScopesItem.SERVICE_BROKER)),
            com.azure.core.util.Context.NONE);
    }
}
