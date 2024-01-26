
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ManagedInstanceAdministratorInner;
import com.azure.resourcemanager.sql.models.AdministratorName;
import com.azure.resourcemanager.sql.models.ManagedInstanceAdministratorType;
import java.util.UUID;

/** Samples for ManagedInstanceAdministrators CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceAdministratorCreate.
     * json
     */
    /**
     * Sample code: Create administrator of managed instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAdministratorOfManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceAdministrators().createOrUpdate(
            "Default-SQL-SouthEastAsia", "managedInstance", AdministratorName.ACTIVE_DIRECTORY,
            new ManagedInstanceAdministratorInner()
                .withAdministratorType(ManagedInstanceAdministratorType.ACTIVE_DIRECTORY).withLogin("bob@contoso.com")
                .withSid(UUID.fromString("44444444-3333-2222-1111-000000000000"))
                .withTenantId(UUID.fromString("55555555-4444-3333-2222-111111111111")),
            Context.NONE);
    }
}
