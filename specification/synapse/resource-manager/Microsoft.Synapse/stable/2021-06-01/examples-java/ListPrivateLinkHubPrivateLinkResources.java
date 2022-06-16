import com.azure.core.util.Context;

/** Samples for PrivateLinkHubPrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListPrivateLinkHubPrivateLinkResources.json
     */
    /**
     * Sample code: Get private link resources for private link hub.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getPrivateLinkResourcesForPrivateLinkHub(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .privateLinkHubPrivateLinkResources()
            .list("ExampleResourceGroup", "ExamplePrivateLinkHub", Context.NONE);
    }
}
