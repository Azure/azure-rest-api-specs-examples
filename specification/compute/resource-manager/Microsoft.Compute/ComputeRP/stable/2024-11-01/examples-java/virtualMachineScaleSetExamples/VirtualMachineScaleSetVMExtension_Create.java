
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetVMExtensionInner;
import java.io.IOException;

/**
 * Samples for VirtualMachineScaleSetVMExtensions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Create.json
     */
    /**
     * Sample code: Create VirtualMachineScaleSet VM extension.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createVirtualMachineScaleSetVMExtension(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMExtensions().createOrUpdate(
            "myResourceGroup", "myvmScaleSet", "0", "myVMExtension",
            new VirtualMachineScaleSetVMExtensionInner().withPublisher("extPublisher").withTypePropertiesType("extType")
                .withTypeHandlerVersion("1.2").withAutoUpgradeMinorVersion(true)
                .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter()
                    .deserialize("{\"UserName\":\"xyz@microsoft.com\"}", Object.class, SerializerEncoding.JSON)),
            com.azure.core.util.Context.NONE);
    }
}
