Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.3/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Databases ListPrincipals. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDatabaseListPrincipals.json
     */
    /**
     * Sample code: KustoDatabaseListPrincipals.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabaseListPrincipals(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.databases().listPrincipals("kustorptest", "kustoclusterrptest4", "KustoDatabase8", Context.NONE);
    }
}
```
