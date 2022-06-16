import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnectionsPrivateLinkHub Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/PrivateEndpointConnectionsPrivateLinkHub_Get.json
     */
    /**
     * Sample code: Get a privateLinkHub.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAPrivateLinkHub(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .privateEndpointConnectionsPrivateLinkHubs()
            .getWithResponse("gh-res-grp", "pe0", "pe0-f3ed30f5-338c-4855-a542-24a403694ad2", Context.NONE);
    }
}
