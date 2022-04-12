Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.ImageInner;

/** Samples for Images CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateAnImageFromAVM.json
     */
    /**
     * Sample code: Create a virtual machine image from an existing virtual machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVirtualMachineImageFromAnExistingVirtualMachine(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getImages()
            .createOrUpdate(
                "myResourceGroup",
                "myImage",
                new ImageInner()
                    .withLocation("West US")
                    .withSourceVirtualMachine(
                        new SubResource()
                            .withId(
                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM")),
                Context.NONE);
    }
}
```
