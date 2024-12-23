
/**
 * Samples for PrivateLinkHubPrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * GetPrivateLinkHubPrivateLinkResource.json
     */
    /**
     * Sample code: Get private link resources for private link hub.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        getPrivateLinkResourcesForPrivateLinkHub(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.privateLinkHubPrivateLinkResources().getWithResponse("ExampleResourceGroup", "ExamplePrivateLinkHub",
            "sql", com.azure.core.util.Context.NONE);
    }
}
