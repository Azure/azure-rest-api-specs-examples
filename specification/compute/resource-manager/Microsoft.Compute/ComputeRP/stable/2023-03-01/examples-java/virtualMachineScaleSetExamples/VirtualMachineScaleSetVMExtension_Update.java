import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMExtensionUpdate;
import java.io.IOException;

/** Samples for VirtualMachineScaleSetVMExtensions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Update.json
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
                com.azure.core.util.Context.NONE);
    }
}
