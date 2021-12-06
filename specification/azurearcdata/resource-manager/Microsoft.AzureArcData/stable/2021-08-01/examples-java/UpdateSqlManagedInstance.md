Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-azurearcdata_1.0.0-beta.2/sdk/azurearcdata/azure-resourcemanager-azurearcdata/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.azurearcdata.models.SqlManagedInstance;
import java.util.HashMap;
import java.util.Map;

/** Samples for SqlManagedInstances Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/UpdateSqlManagedInstance.json
     */
    /**
     * Sample code: Updates a sql Instance tags.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void updatesASqlInstanceTags(com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        SqlManagedInstance resource =
            manager
                .sqlManagedInstances()
                .getByResourceGroupWithResponse("testrg", "testsqlManagedInstance", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("mytag", "myval")).apply();
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
