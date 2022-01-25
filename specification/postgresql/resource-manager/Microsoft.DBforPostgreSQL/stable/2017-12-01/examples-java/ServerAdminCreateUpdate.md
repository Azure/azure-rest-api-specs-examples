Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.postgresql.fluent.models.ServerAdministratorResourceInner;
import com.azure.resourcemanager.postgresql.models.AdministratorType;
import java.util.UUID;

/** Samples for ServerAdministrators CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerAdminCreateUpdate.json
     */
    /**
     * Sample code: ServerAdministratorCreate.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverAdministratorCreate(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .serverAdministrators()
            .createOrUpdate(
                "testrg",
                "pgtestsvc4",
                new ServerAdministratorResourceInner()
                    .withAdministratorType(AdministratorType.ACTIVE_DIRECTORY)
                    .withLogin("bob@contoso.com")
                    .withSid(UUID.fromString("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"))
                    .withTenantId(UUID.fromString("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c")),
                Context.NONE);
    }
}
```
