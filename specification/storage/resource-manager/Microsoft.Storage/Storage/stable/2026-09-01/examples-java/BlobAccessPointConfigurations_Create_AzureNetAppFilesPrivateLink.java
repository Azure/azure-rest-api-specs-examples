
import com.azure.resourcemanager.storage.fluent.models.BlobAccessPointConfigurationInner;
import com.azure.resourcemanager.storage.models.BlobAccessPointAccessKeyAuthProperties;
import com.azure.resourcemanager.storage.models.BlobAccessPointAzureNetAppFilesSourceProperties;
import com.azure.resourcemanager.storage.models.BlobAccessPointConfigurationProperties;
import com.azure.resourcemanager.storage.models.BlobAccessPointConfigurationState;
import com.azure.resourcemanager.storage.models.BlobAccessPointPrivateLinkConnectionProperties;
import com.azure.resourcemanager.storage.models.BlobAccessPointPrivateLinkIdType;
import com.azure.resourcemanager.storage.models.BlobAccessPointTlsVerification;

/**
 * Samples for BlobAccessPointConfigurations Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/BlobAccessPointConfigurations_Create_AzureNetAppFilesPrivateLink.json
     */
    /**
     * Sample code: CreateBlobAccessPointConfigurationWithAzureNetAppFilesPrivateLink.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void createBlobAccessPointConfigurationWithAzureNetAppFilesPrivateLink(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobAccessPointConfigurations().create("testrg", "teststorageaccount",
            "anfprivatelinkaccesspoint",
            new BlobAccessPointConfigurationInner().withLocation("eastus")
                .withProperties(new BlobAccessPointConfigurationProperties()
                    .withState(BlobAccessPointConfigurationState.ACTIVE)
                    .withDescription("Azure NetApp Files S3-compatible source over Private Link")
                    .withSource(new BlobAccessPointAzureNetAppFilesSourceProperties()
                        .withConnection(new BlobAccessPointPrivateLinkConnectionProperties()
                            .withPrivateLinkIdType(BlobAccessPointPrivateLinkIdType.RESOURCE_ID)
                            .withPrivateLinkId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/privateLinkServices/anfpls")
                            .withPrivateLinkGroupId("s3").withPrivateLinkLocation("eastus")
                            .withRequestMessage("Blob Access Point connection request for Azure NetApp Files")
                            .withEndpoint("https://anf.contoso.com/bucket1")
                            .withTlsVerification(BlobAccessPointTlsVerification.PERFORM))
                        .withAuth(new BlobAccessPointAccessKeyAuthProperties().withAccessKeyId("fakeTokenPlaceholder")
                            .withSecretAccessKey("fakeTokenPlaceholder").withSigningRegion("us-east-1")))),
            com.azure.core.util.Context.NONE);
    }
}
