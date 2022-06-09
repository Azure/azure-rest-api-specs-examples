```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Backup;
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateModeForUpdate;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Server;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Sku;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SkuTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Storage;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/ServerUpdate.json
     */
    /**
     * Sample code: ServerUpdate.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverUpdate(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource =
            manager.servers().getByResourceGroupWithResponse("TestGroup", "pgtestsvc4", Context.NONE).getValue();
        resource
            .update()
            .withSku(new Sku().withName("Standard_D8s_v3").withTier(SkuTier.GENERAL_PURPOSE))
            .withAdministratorLoginPassword("newpassword")
            .withStorage(new Storage().withStorageSizeGB(1024))
            .withBackup(new Backup().withBackupRetentionDays(20))
            .withCreateMode(CreateModeForUpdate.UPDATE)
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresqlflexibleserver_1.0.0-beta.4/sdk/postgresqlflexibleserver/azure-resourcemanager-postgresqlflexibleserver/README.md) on how to add the SDK to your project and authenticate.
