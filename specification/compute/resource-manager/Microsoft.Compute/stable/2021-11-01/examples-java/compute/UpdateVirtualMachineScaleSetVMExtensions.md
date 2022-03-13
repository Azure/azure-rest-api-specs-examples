Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMExtensionUpdate;
import java.io.IOException;

/** Samples for VirtualMachineScaleSetVMExtensions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/UpdateVirtualMachineScaleSetVMExtensions.json
     */
    /**
     * Sample code: Update VirtualMachineScaleSet VM extension.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateVirtualMachineScaleSetVMExtension(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMExtensions()
            .update(
                "myResourceGroup",
                "myvmScaleSet",
                "0",
                "myVMExtension",
                new VirtualMachineScaleSetVMExtensionUpdate()
                    .withPublisher("extPublisher")
                    .withTypePropertiesType("extType")
                    .withTypeHandlerVersion("1.2")
                    .withAutoUpgradeMinorVersion(true)
                    .withSettings(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize(
                                "{\"UserName\":\"xyz@microsoft.com\"}", Object.class, SerializerEncoding.JSON)),
                Context.NONE);
    }
}
```
