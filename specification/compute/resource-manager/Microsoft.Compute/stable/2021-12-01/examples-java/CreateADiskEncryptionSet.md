Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.DiskEncryptionSetInner;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetIdentityType;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetType;
import com.azure.resourcemanager.compute.models.EncryptionSetIdentity;
import com.azure.resourcemanager.compute.models.KeyForDiskEncryptionSet;
import com.azure.resourcemanager.compute.models.SourceVault;

/** Samples for DiskEncryptionSets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/CreateADiskEncryptionSet.json
     */
    /**
     * Sample code: Create a disk encryption set.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADiskEncryptionSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskEncryptionSets()
            .createOrUpdate(
                "myResourceGroup",
                "myDiskEncryptionSet",
                new DiskEncryptionSetInner()
                    .withLocation("West US")
                    .withIdentity(new EncryptionSetIdentity().withType(DiskEncryptionSetIdentityType.SYSTEM_ASSIGNED))
                    .withEncryptionType(DiskEncryptionSetType.ENCRYPTION_AT_REST_WITH_CUSTOMER_KEY)
                    .withActiveKey(
                        new KeyForDiskEncryptionSet()
                            .withSourceVault(
                                new SourceVault()
                                    .withId(
                                        "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"))
                            .withKeyUrl("https://myvmvault.vault-int.azure-int.net/keys/{key}")),
                Context.NONE);
    }
}
```
