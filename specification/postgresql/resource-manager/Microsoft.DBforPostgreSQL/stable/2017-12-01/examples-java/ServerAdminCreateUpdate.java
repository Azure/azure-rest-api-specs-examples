
import com.azure.resourcemanager.postgresql.fluent.models.ServerAdministratorResourceInner;
import com.azure.resourcemanager.postgresql.models.AdministratorType;
import java.util.UUID;

/**
 * Samples for ServerAdministrators CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/
     * ServerAdminCreateUpdate.json
     */
    /**
     * Sample code: ServerAdministratorCreate.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverAdministratorCreate(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.serverAdministrators().createOrUpdate("testrg", "pgtestsvc4",
            new ServerAdministratorResourceInner().withAdministratorType(AdministratorType.ACTIVE_DIRECTORY)
                .withLogin("bob@contoso.com").withSid(UUID.fromString("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"))
                .withTenantId(UUID.fromString("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c")),
            com.azure.core.util.Context.NONE);
    }
}
