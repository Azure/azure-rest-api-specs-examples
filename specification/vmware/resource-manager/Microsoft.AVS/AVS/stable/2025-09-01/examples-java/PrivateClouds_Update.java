
import com.azure.resourcemanager.avs.models.Encryption;
import com.azure.resourcemanager.avs.models.EncryptionKeyVaultProperties;
import com.azure.resourcemanager.avs.models.EncryptionState;
import com.azure.resourcemanager.avs.models.ManagementCluster;
import com.azure.resourcemanager.avs.models.PrivateCloud;
import com.azure.resourcemanager.avs.models.PrivateCloudIdentity;
import com.azure.resourcemanager.avs.models.ResourceIdentityType;

/**
 * Samples for PrivateClouds Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PrivateClouds_Update.json
     */
    /**
     * Sample code: PrivateClouds_Update.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        PrivateCloud resource = manager.privateClouds()
            .getByResourceGroupWithResponse("group1", "cloud1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withIdentity(new PrivateCloudIdentity().withType(ResourceIdentityType.NONE))
            .withManagementCluster(new ManagementCluster().withClusterSize(4))
            .withEncryption(new Encryption().withStatus(EncryptionState.ENABLED)
                .withKeyVaultProperties(new EncryptionKeyVaultProperties().withKeyName("fakeTokenPlaceholder")
                    .withKeyVersion("fakeTokenPlaceholder").withKeyVaultUrl("fakeTokenPlaceholder")))
            .apply();
    }
}
