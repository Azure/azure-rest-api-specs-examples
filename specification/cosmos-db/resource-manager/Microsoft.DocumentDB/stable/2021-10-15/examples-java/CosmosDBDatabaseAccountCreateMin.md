Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.CreateMode;
import com.azure.resourcemanager.cosmos.models.DatabaseAccountCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.Location;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for DatabaseAccounts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBDatabaseAccountCreateMin.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountCreateMin.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountCreateMin(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getDatabaseAccounts()
            .createOrUpdate(
                "rg1",
                "ddb1",
                new DatabaseAccountCreateUpdateParameters()
                    .withLocation("westus")
                    .withLocations(
                        Arrays
                            .asList(
                                new Location()
                                    .withLocationName("southcentralus")
                                    .withFailoverPriority(0)
                                    .withIsZoneRedundant(false)))
                    .withCreateMode(CreateMode.DEFAULT),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```
