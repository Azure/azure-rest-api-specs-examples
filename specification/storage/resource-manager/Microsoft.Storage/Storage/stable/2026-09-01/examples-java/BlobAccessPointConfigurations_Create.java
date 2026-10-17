
import com.azure.resourcemanager.storage.fluent.models.BlobAccessPointConfigurationInner;
import com.azure.resourcemanager.storage.models.BlobAccessPointAccessKeyAuthProperties;
import com.azure.resourcemanager.storage.models.BlobAccessPointConfigurationProperties;
import com.azure.resourcemanager.storage.models.BlobAccessPointConfigurationState;
import com.azure.resourcemanager.storage.models.BlobAccessPointEndpointConnectionProperties;
import com.azure.resourcemanager.storage.models.BlobAccessPointNetAppOntapSourceProperties;
import com.azure.resourcemanager.storage.models.BlobAccessPointTlsVerification;

/**
 * Samples for BlobAccessPointConfigurations Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/BlobAccessPointConfigurations_Create.json
     */
    /**
     * Sample code: CreateBlobAccessPointConfiguration.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void createBlobAccessPointConfiguration(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobAccessPointConfigurations().create("testrg", "teststorageaccount",
            "testaccesspointconfig",
            new BlobAccessPointConfigurationInner().withLocation("eastus")
                .withProperties(new BlobAccessPointConfigurationProperties()
                    .withState(BlobAccessPointConfigurationState.ACTIVE)
                    .withDescription("NetApp ONTAP S3-compatible source")
                    .withSource(new BlobAccessPointNetAppOntapSourceProperties()
                        .withConnection(new BlobAccessPointEndpointConnectionProperties()
                            .withEndpoint("https://netapp.contoso.com/bucket1")
                            .withTlsVerification(BlobAccessPointTlsVerification.PERFORM))
                        .withAuth(new BlobAccessPointAccessKeyAuthProperties().withAccessKeyId("fakeTokenPlaceholder")
                            .withSecretAccessKey("fakeTokenPlaceholder").withSigningRegion("us-east-1")))),
            com.azure.core.util.Context.NONE);
    }
}
