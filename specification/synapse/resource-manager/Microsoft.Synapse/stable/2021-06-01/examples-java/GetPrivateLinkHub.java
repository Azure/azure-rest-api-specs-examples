/** Samples for PrivateLinkHubs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetPrivateLinkHub.json
     */
    /**
     * Sample code: Get a privateLinkHub.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAPrivateLinkHub(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .privateLinkHubs()
            .getByResourceGroupWithResponse("resourceGroup1", "privateLinkHub1", com.azure.core.util.Context.NONE);
    }
}
