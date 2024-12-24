
/**
 * Samples for Workspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/ListWorkspacesInSubscription.
     * json
     */
    /**
     * Sample code: List workspaces in subscription.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void listWorkspacesInSubscription(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaces().list(com.azure.core.util.Context.NONE);
    }
}
