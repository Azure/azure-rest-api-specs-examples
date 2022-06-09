```java
import com.azure.core.util.Context;

/** Samples for VirtualMachineImageTemplates Run. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/RunImageTemplate.json
     */
    /**
     * Sample code: Create image(s) from existing imageTemplate.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void createImageSFromExistingImageTemplate(
        com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().run("myResourceGroup", "myImageTemplate", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-imagebuilder_1.0.0-beta.2/sdk/imagebuilder/azure-resourcemanager-imagebuilder/README.md) on how to add the SDK to your project and authenticate.
