```java
import com.azure.core.util.Context;

/** Samples for Clusters ListSkusByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClustersListResourceSkus.json
     */
    /**
     * Sample code: KustoClustersListResourceSkus.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersListResourceSkus(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().listSkusByResource("kustorptest", "kustoCluster", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.
