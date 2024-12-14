
import com.azure.resourcemanager.hybridnetwork.models.ArtifactReplicationStrategy;
import com.azure.resourcemanager.hybridnetwork.models.ArtifactStorePropertiesFormat;
import com.azure.resourcemanager.hybridnetwork.models.ArtifactStorePropertiesFormatManagedResourceGroupConfiguration;
import com.azure.resourcemanager.hybridnetwork.models.ArtifactStoreType;

/**
 * Samples for ArtifactStores CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ArtifactStoreCreate.json
     */
    /**
     * Sample code: Create or update an artifact store of publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdateAnArtifactStoreOfPublisherResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.artifactStores().define("TestArtifactStore").withRegion("eastus")
            .withExistingPublisher("rg", "TestPublisher")
            .withProperties(
                new ArtifactStorePropertiesFormat().withStoreType(ArtifactStoreType.AZURE_CONTAINER_REGISTRY)
                    .withReplicationStrategy(ArtifactReplicationStrategy.SINGLE_REPLICATION)
                    .withManagedResourceGroupConfiguration(
                        new ArtifactStorePropertiesFormatManagedResourceGroupConfiguration().withName("testRg")
                            .withLocation("eastus")))
            .create();
    }
}
