Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.DatabaseAccountRegenerateKeyParameters;
import com.azure.resourcemanager.cosmos.models.KeyKind;

/** Samples for DatabaseAccounts RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBDatabaseAccountRegenerateKey.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountRegenerateKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountRegenerateKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getDatabaseAccounts()
            .regenerateKey(
                "rg1", "ddb1", new DatabaseAccountRegenerateKeyParameters().withKeyKind(KeyKind.PRIMARY), Context.NONE);
    }
}
```
