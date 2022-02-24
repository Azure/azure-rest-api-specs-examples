Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.Script;

/** Samples for Scripts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoScriptsUpdate.json
     */
    /**
     * Sample code: KustoScriptsUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        Script resource =
            manager
                .scripts()
                .getWithResponse("kustorptest", "kustoCluster", "KustoDatabase8", "kustoScript", Context.NONE)
                .getValue();
        resource
            .update()
            .withScriptUrl("https://mysa.blob.core.windows.net/container/script.txt")
            .withForceUpdateTag("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe")
            .withContinueOnErrors(true)
            .apply();
    }
}
```
