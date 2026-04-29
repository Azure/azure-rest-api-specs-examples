
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetVMExtensionInner;
import java.io.IOException;

/**
 * Samples for VirtualMachineScaleSetVMExtensions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Create.json
     */
    /**
     * Sample code: Create VirtualMachineScaleSet VM extension.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createVirtualMachineScaleSetVMExtension(com.azure.resourcemanager.compute.ComputeManager manager)
        throws IOException {
        manager.serviceClient().getVirtualMachineScaleSetVMExtensions().createOrUpdate("myResourceGroup",
            "myvmScaleSet", "0", "myVMExtension",
            new VirtualMachineScaleSetVMExtensionInner().withPublisher("extPublisher").withTypePropertiesType("extType")
                .withTypeHandlerVersion("1.2").withAutoUpgradeMinorVersion(true)
                .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter()
                    .deserialize("{\"UserName\":\"xyz@microsoft.com\"}", Object.class, SerializerEncoding.JSON)),
            com.azure.core.util.Context.NONE);
    }
}
