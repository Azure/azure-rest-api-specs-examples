Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.ContinuousBackupRestoreLocation;

/** Samples for SqlResources RetrieveContinuousBackupInformation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlContainerBackupInformation.json
     */
    /**
     * Sample code: CosmosDBSqlContainerBackupInformation.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlContainerBackupInformation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .retrieveContinuousBackupInformation(
                "rgName",
                "ddb1",
                "databaseName",
                "containerName",
                new ContinuousBackupRestoreLocation().withLocation("North Europe"),
                Context.NONE);
    }
}
```
