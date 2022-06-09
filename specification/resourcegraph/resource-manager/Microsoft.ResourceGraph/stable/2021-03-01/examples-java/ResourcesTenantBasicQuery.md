```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.resourcegraph.models.QueryRequest;

/** Samples for ResourceProvider Resources. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2021-03-01/examples/ResourcesTenantBasicQuery.json
     */
    /**
     * Sample code: Basic tenant query.
     *
     * @param manager Entry point to ResourceGraphManager.
     */
    public static void basicTenantQuery(com.azure.resourcemanager.resourcegraph.ResourceGraphManager manager) {
        manager
            .resourceProviders()
            .resourcesWithResponse(
                new QueryRequest().withQuery("Resources | project id, name, type, location, tags | limit 3"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-resourcegraph_1.0.0-beta.3/sdk/resourcegraph/azure-resourcemanager-resourcegraph/README.md) on how to add the SDK to your project and authenticate.
