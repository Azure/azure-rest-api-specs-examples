import com.azure.core.util.Context;
import com.azure.resourcemanager.mysql.fluent.models.ServerAdministratorResourceInner;
import com.azure.resourcemanager.mysql.models.AdministratorType;
import java.util.UUID;

/** Samples for ServerAdministrators CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerAdminCreateUpdate.json
     */
    /**
     * Sample code: ServerAdministratorCreate.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverAdministratorCreate(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .serverAdministrators()
            .createOrUpdate(
                "testrg",
                "mysqltestsvc4",
                new ServerAdministratorResourceInner()
                    .withAdministratorType(AdministratorType.ACTIVE_DIRECTORY)
                    .withLogin("bob@contoso.com")
                    .withSid(UUID.fromString("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"))
                    .withTenantId(UUID.fromString("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c")),
                Context.NONE);
    }
}
