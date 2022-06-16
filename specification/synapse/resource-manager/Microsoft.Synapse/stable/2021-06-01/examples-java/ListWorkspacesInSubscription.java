import com.azure.core.util.Context;

/** Samples for Workspaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspacesInSubscription.json
     */
    /**
     * Sample code: List workspaces in subscription.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listWorkspacesInSubscription(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaces().list(Context.NONE);
    }
}
