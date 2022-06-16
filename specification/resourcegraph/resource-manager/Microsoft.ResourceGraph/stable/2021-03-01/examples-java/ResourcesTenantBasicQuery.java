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
