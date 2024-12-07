
import com.azure.resourcemanager.resourcegraph.models.QueryRequest;
import com.azure.resourcemanager.resourcegraph.models.QueryRequestOptions;
import java.util.Arrays;

/**
 * Samples for ResourceProvider Resources.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2021-03-01/examples/
     * ResourcesNextPageQuery.json
     */
    /**
     * Sample code: Next page query.
     * 
     * @param manager Entry point to ResourceGraphManager.
     */
    public static void nextPageQuery(com.azure.resourcemanager.resourcegraph.ResourceGraphManager manager) {
        manager.resourceProviders()
            .resourcesWithResponse(
                new QueryRequest().withSubscriptions(Arrays.asList("cfbbd179-59d2-4052-aa06-9270a38aa9d6"))
                    .withQuery("Resources | where name contains 'test' | project id, name, type, location")
                    .withOptions(new QueryRequestOptions().withSkipToken("fakeTokenPlaceholder")),
                com.azure.core.util.Context.NONE);
    }
}
