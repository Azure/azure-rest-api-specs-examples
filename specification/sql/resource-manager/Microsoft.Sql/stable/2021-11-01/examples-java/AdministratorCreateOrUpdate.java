
import com.azure.resourcemanager.sql.fluent.models.ServerAzureADAdministratorInner;
import com.azure.resourcemanager.sql.models.AdministratorName;
import com.azure.resourcemanager.sql.models.AdministratorType;
import java.util.UUID;

/**
 * Samples for ServerAzureADAdministrators CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/AdministratorCreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates an existing Azure Active Directory administrator.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsOrUpdatesAnExistingAzureActiveDirectoryAdministrator(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAzureADAdministrators().createOrUpdate("sqlcrudtest-4799",
            "sqlcrudtest-6440", AdministratorName.ACTIVE_DIRECTORY,
            new ServerAzureADAdministratorInner().withAdministratorType(AdministratorType.ACTIVE_DIRECTORY)
                .withLogin("bob@contoso.com").withSid(UUID.fromString("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"))
                .withTenantId(UUID.fromString("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c")),
            com.azure.core.util.Context.NONE);
    }
}
