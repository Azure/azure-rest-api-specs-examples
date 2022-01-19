Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for RestorableSqlContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBRestorableSqlContainerList.json
     */
    /**
     * Sample code: CosmosDBRestorableSqlContainerList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBRestorableSqlContainerList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getRestorableSqlContainers()
            .list("WestUS", "98a570f2-63db-4117-91f0-366327b7b353", "3fu-hg==", null, null, Context.NONE);
    }
}
```
