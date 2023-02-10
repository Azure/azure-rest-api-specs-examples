/** Samples for PrivateLinkHubs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeletePrivateLinkHub.json
     */
    /**
     * Sample code: Delete a privateLinkHub.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteAPrivateLinkHub(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.privateLinkHubs().delete("resourceGroup1", "privateLinkHub1", com.azure.core.util.Context.NONE);
    }
}
