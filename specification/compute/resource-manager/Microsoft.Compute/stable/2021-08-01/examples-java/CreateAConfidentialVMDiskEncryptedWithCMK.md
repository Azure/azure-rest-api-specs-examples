```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.DiskSecurityProfile;
import com.azure.resourcemanager.compute.models.DiskSecurityTypes;
import com.azure.resourcemanager.compute.models.ImageDiskReference;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/** Samples for Disks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/CreateAConfidentialVMDiskEncryptedWithCMK.json
     */
    /**
     * Sample code: Create a confidential VM supported disk encrypted with customer managed key.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAConfidentialVMSupportedDiskEncryptedWithCustomerManagedKey(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .createOrUpdate(
                "myResourceGroup",
                "myDisk",
                new DiskInner()
                    .withLocation("West US")
                    .withOsType(OperatingSystemTypes.WINDOWS)
                    .withCreationData(
                        new CreationData()
                            .withCreateOption(DiskCreateOption.FROM_IMAGE)
                            .withImageReference(
                                new ImageDiskReference()
                                    .withId(
                                        "/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/westus/Publishers/{publisher}/ArtifactTypes/VMImage/Offers/{offer}/Skus/{sku}/Versions/1.0.0")))
                    .withSecurityProfile(
                        new DiskSecurityProfile()
                            .withSecurityType(DiskSecurityTypes.CONFIDENTIAL_VM_DISK_ENCRYPTED_WITH_CUSTOMER_KEY)
                            .withSecureVMDiskEncryptionSetId(
                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
