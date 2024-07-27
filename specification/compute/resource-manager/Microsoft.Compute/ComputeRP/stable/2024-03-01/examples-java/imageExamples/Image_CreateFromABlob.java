
import com.azure.resourcemanager.compute.fluent.models.ImageInner;
import com.azure.resourcemanager.compute.models.ImageOSDisk;
import com.azure.resourcemanager.compute.models.ImageStorageProfile;
import com.azure.resourcemanager.compute.models.OperatingSystemStateTypes;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/**
 * Samples for Images CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/imageExamples/
     * Image_CreateFromABlob.json
     */
    /**
     * Sample code: Create a virtual machine image from a blob.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVirtualMachineImageFromABlob(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getImages().createOrUpdate("myResourceGroup", "myImage",
            new ImageInner().withLocation("West US").withStorageProfile(new ImageStorageProfile()
                .withOsDisk(
                    new ImageOSDisk().withBlobUri("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd")
                        .withOsType(OperatingSystemTypes.LINUX).withOsState(OperatingSystemStateTypes.GENERALIZED))
                .withZoneResilient(true)),
            com.azure.core.util.Context.NONE);
    }
}
