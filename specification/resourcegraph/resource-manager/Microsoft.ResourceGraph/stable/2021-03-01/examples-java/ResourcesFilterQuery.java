import com.azure.core.util.Context;
import com.azure.resourcemanager.resourcegraph.models.QueryRequest;
import java.util.Arrays;

/** Samples for ResourceProvider Resources. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2021-03-01/examples/ResourcesFilterQuery.json
     */
    /**
     * Sample code: Filter resources.
     *
     * @param manager Entry point to ResourceGraphManager.
     */
    public static void filterResources(com.azure.resourcemanager.resourcegraph.ResourceGraphManager manager) {
        manager
            .resourceProviders()
            .resourcesWithResponse(
                new QueryRequest()
                    .withSubscriptions(Arrays.asList("cfbbd179-59d2-4052-aa06-9270a38aa9d6"))
                    .withQuery(
                        "Resources | project id, name, type, location | where type =~"
                            + " 'Microsoft.Compute/virtualMachines' | limit 3"),
                Context.NONE);
    }
}
