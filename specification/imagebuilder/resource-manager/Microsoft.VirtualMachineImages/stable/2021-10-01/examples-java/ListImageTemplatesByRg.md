Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-imagebuilder_1.0.0-beta.2/sdk/imagebuilder/azure-resourcemanager-imagebuilder/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VirtualMachineImageTemplates ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/ListImageTemplatesByRg.json
     */
    /**
     * Sample code: List images by resource group.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void listImagesByResourceGroup(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
```
