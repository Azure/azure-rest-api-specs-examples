```java
import com.azure.core.util.Context;

/** Samples for ClusterPrincipalAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClusterPrincipalAssignmentsDelete.json
     */
    /**
     * Sample code: KustoClusterPrincipalAssignmentsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterPrincipalAssignmentsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusterPrincipalAssignments().delete("kustorptest", "kustoCluster", "kustoprincipal1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.
