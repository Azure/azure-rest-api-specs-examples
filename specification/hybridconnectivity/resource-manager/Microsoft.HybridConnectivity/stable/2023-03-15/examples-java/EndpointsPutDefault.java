import com.azure.resourcemanager.hybridconnectivity.models.EndpointProperties;
import com.azure.resourcemanager.hybridconnectivity.models.Type;

/** Samples for Endpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/EndpointsPutDefault.json
     */
    /**
     * Sample code: HybridConnectivityEndpointsPutDefault.
     *
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityEndpointsPutDefault(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager
            .endpoints()
            .define("default")
            .withExistingResourceUri(
                "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine")
            .withProperties(new EndpointProperties().withType(Type.DEFAULT))
            .create();
    }
}
