import com.azure.resourcemanager.hybridconnectivity.models.EndpointProperties;
import com.azure.resourcemanager.hybridconnectivity.models.Type;

/** Samples for Endpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/EndpointsPutCustom.json
     */
    /**
     * Sample code: HybridConnectivityEndpointsPutCustom.
     *
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityEndpointsPutCustom(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager
            .endpoints()
            .define("custom")
            .withExistingResourceUri(
                "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine")
            .withProperties(
                new EndpointProperties()
                    .withType(Type.CUSTOM)
                    .withResourceId(
                        "/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.Relay/namespaces/custom-relay-namespace"))
            .create();
    }
}
