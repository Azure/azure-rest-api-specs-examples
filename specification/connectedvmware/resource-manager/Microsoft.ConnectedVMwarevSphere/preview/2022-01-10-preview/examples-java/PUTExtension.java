import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/** Samples for MachineExtensions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/PUTExtension.json
     */
    /**
     * Sample code: Create or Update a Machine Extension (PUT).
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void createOrUpdateAMachineExtensionPUT(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) throws IOException {
        manager
            .machineExtensions()
            .define("CustomScriptExtension")
            .withRegion("eastus2euap")
            .withExistingVirtualMachine("myResourceGroup", "myMachine")
            .withPublisher("Microsoft.Compute")
            .withTypePropertiesType("CustomScriptExtension")
            .withTypeHandlerVersion("1.10")
            .withSettings(
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize(
                        "{\"commandToExecute\":\"powershell.exe -c \\\"Get-Process | Where-Object { $_.CPU -gt 10000"
                            + " }\\\"\"}",
                        Object.class,
                        SerializerEncoding.JSON))
            .create();
    }
}
