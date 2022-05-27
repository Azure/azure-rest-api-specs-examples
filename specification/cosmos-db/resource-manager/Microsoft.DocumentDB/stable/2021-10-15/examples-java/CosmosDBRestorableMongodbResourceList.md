Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for RestorableMongodbResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBRestorableMongodbResourceList.json
     */
    /**
     * Sample code: CosmosDBRestorableMongodbResourceList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBRestorableMongodbResourceList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getRestorableMongodbResources()
            .list("WestUS", "d9b26648-2f53-4541-b3d8-3044f4f9810d", "WestUS", "10/13/2020 4:56", Context.NONE);
    }
}
```
