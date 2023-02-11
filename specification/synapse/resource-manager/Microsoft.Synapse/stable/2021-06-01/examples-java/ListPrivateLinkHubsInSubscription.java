/** Samples for PrivateLinkHubs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListPrivateLinkHubsInSubscription.json
     */
    /**
     * Sample code: List privateLinkHubs in subscription.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listPrivateLinkHubsInSubscription(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.privateLinkHubs().list(com.azure.core.util.Context.NONE);
    }
}
