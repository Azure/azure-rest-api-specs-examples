```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.ScriptCheckNameRequest;

/** Samples for Scripts CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoScriptsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoScriptsCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsCheckNameAvailability(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .scripts()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoCluster",
                "db",
                new ScriptCheckNameRequest().withName("kustoScriptName1"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.
