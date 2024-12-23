
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/GetPrivateLinkResource.json
     */
    /**
     * Sample code: Get private link resources for workspace.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getPrivateLinkResourcesForWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.privateLinkResources().getWithResponse("ExampleResourceGroup", "ExampleWorkspace", "sql",
            com.azure.core.util.Context.NONE);
    }
}
