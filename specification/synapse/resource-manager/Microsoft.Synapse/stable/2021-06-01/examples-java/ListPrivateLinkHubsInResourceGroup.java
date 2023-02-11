/** Samples for PrivateLinkHubs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListPrivateLinkHubsInResourceGroup.json
     */
    /**
     * Sample code: List privateLinkHubs in resource group.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listPrivateLinkHubsInResourceGroup(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.privateLinkHubs().listByResourceGroup("resourceGroup1", com.azure.core.util.Context.NONE);
    }
}
