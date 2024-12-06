
/**
 * Samples for GraphQuery GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2021-03-01/examples/GraphQueryGet.
     * json
     */
    /**
     * Sample code: Get Graph Query.
     * 
     * @param manager Entry point to ResourceGraphManager.
     */
    public static void getGraphQuery(com.azure.resourcemanager.resourcegraph.ResourceGraphManager manager) {
        manager.graphQueries().getByResourceGroupWithResponse("024e2271-06fa-46b6-9079-f1ed3c7b070e",
            "my-resource-group", "MyDockerVMs", com.azure.core.util.Context.NONE);
    }
}
