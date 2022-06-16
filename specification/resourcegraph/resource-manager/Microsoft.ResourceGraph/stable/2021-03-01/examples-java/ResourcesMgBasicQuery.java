import com.azure.core.util.Context;
import com.azure.resourcemanager.resourcegraph.models.QueryRequest;
import java.util.Arrays;

/** Samples for ResourceProvider Resources. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2021-03-01/examples/ResourcesMgBasicQuery.json
     */
    /**
     * Sample code: Basic management group query.
     *
     * @param manager Entry point to ResourceGraphManager.
     */
    public static void basicManagementGroupQuery(com.azure.resourcemanager.resourcegraph.ResourceGraphManager manager) {
        manager
            .resourceProviders()
            .resourcesWithResponse(
                new QueryRequest()
                    .withManagementGroups(Arrays.asList("e927f598-c1d4-4f72-8541-95d83a6a4ac8", "ProductionMG"))
                    .withQuery("Resources | project id, name, type, location, tags | limit 3"),
                Context.NONE);
    }
}
