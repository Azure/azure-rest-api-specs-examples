Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.3/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for Scripts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoScriptsCreateOrUpdate.json
     */
    /**
     * Sample code: KustoScriptsCreateOrUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsCreateOrUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .scripts()
            .define("kustoScript1")
            .withExistingDatabase("kustorptest", "kustoclusterrptest4", "KustoDatabase8")
            .withScriptUrl("https://mysa.blob.core.windows.net/container/script.txt")
            .withScriptUrlSasToken(
                "?sv=2019-02-02&st=2019-04-29T22%3A18%3A26Z&se=2019-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=********************************")
            .withForceUpdateTag("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe")
            .withContinueOnErrors(true)
            .create();
    }
}
```
