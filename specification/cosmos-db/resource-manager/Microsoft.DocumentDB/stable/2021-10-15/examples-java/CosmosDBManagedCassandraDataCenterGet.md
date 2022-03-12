Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for CassandraDataCenters Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBManagedCassandraDataCenterGet.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraDataCenterGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraDataCenterGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraDataCenters()
            .getWithResponse("cassandra-prod-rg", "cassandra-prod", "dc1", Context.NONE);
    }
}
```
