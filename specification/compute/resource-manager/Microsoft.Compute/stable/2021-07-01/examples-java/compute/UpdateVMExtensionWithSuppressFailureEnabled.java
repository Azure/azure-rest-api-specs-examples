import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.compute.models.VirtualMachineExtensionUpdate;
import java.io.IOException;

/** Samples for VirtualMachineExtensions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/UpdateVMExtensionWithSuppressFailureEnabled.json
     */
    /**
     * Sample code: Update VM extension.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateVMExtension(com.azure.resourcemanager.AzureResourceManager azure) throws IOException {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensions()
            .update(
                "myResourceGroup",
                "myVM",
                "myVMExtension",
                new VirtualMachineExtensionUpdate()
                    .withPublisher("extPublisher")
                    .withType("extType")
                    .withTypeHandlerVersion("1.2")
                    .withAutoUpgradeMinorVersion(true)
                    .withSettings(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize("{\"UserName\":\"xyz@microsoft.com\"}", Object.class, SerializerEncoding.JSON))
                    .withSuppressFailures(true),
                Context.NONE);
    }
}
