Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.avs.models.Encryption;
import com.azure.resourcemanager.avs.models.EncryptionKeyVaultProperties;
import com.azure.resourcemanager.avs.models.EncryptionState;
import com.azure.resourcemanager.avs.models.ManagementCluster;
import com.azure.resourcemanager.avs.models.PrivateCloud;
import com.azure.resourcemanager.avs.models.PrivateCloudIdentity;
import com.azure.resourcemanager.avs.models.ResourceIdentityType;

/** Samples for PrivateClouds Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_Update.json
     */
    /**
     * Sample code: PrivateClouds_Update.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        PrivateCloud resource =
            manager.privateClouds().getByResourceGroupWithResponse("group1", "cloud1", Context.NONE).getValue();
        resource
            .update()
            .withIdentity(new PrivateCloudIdentity().withType(ResourceIdentityType.NONE))
            .withManagementCluster(new ManagementCluster().withClusterSize(4))
            .withEncryption(
                new Encryption()
                    .withStatus(EncryptionState.ENABLED)
                    .withKeyVaultProperties(
                        new EncryptionKeyVaultProperties()
                            .withKeyName("keyname1")
                            .withKeyVersion("ver1.0")
                            .withKeyVaultUrl("https://keyvault1-kmip-kvault.vault.azure.net/")))
            .apply();
    }
}
```
