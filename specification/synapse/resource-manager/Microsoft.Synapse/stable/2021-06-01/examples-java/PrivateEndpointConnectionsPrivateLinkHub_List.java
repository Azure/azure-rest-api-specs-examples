/** Samples for PrivateEndpointConnectionsPrivateLinkHub List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/PrivateEndpointConnectionsPrivateLinkHub_List.json
     */
    /**
     * Sample code: Get a privateLinkHub.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAPrivateLinkHub(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.privateEndpointConnectionsPrivateLinkHubs().list("gh-res-grp", "pe0", com.azure.core.util.Context.NONE);
    }
}
