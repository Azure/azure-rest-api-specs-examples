Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.3/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.CheckNameRequest;
import com.azure.resourcemanager.kusto.models.Type;

/** Samples for Databases CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDatabasesCheckNameAvailability.json
     */
    /**
     * Sample code: KustoDatabasesCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesCheckNameAvailability(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databases()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoclusterrptest4",
                new CheckNameRequest().withName("database1").withType(Type.MICROSOFT_KUSTO_CLUSTERS_DATABASES),
                Context.NONE);
    }
}
```
