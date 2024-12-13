
import com.azure.resourcemanager.hybridnetwork.models.PublisherPropertiesFormat;
import com.azure.resourcemanager.hybridnetwork.models.PublisherScope;

/**
 * Samples for Publishers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/PublisherCreate.
     * json
     */
    /**
     * Sample code: Create or update a publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        createOrUpdateAPublisherResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.publishers().define("TestPublisher").withRegion("eastus").withExistingResourceGroup("rg")
            .withProperties(new PublisherPropertiesFormat().withScope(PublisherScope.fromString("Public"))).create();
    }
}
