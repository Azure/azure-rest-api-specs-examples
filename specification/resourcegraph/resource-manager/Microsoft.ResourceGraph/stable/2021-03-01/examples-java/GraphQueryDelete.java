
/**
 * Samples for GraphQuery Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2021-03-01/examples/GraphQueryDelete.
     * json
     */
    /**
     * Sample code: Delete Graph Query.
     * 
     * @param manager Entry point to ResourceGraphManager.
     */
    public static void deleteGraphQuery(com.azure.resourcemanager.resourcegraph.ResourceGraphManager manager) {
        manager.graphQueries().deleteWithResponse("024e2271-06fa-46b6-9079-f1ed3c7b070e", "my-resource-group",
            "MyDockerVM", com.azure.core.util.Context.NONE);
    }
}
