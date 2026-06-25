
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetExtensionUpdate;
import java.io.IOException;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSetExtensions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtension_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetExtensionUpdateMaximumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) throws IOException {
        manager.serviceClient().getVirtualMachineScaleSetExtensions().update("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaa",
            new VirtualMachineScaleSetExtensionUpdate().withForceUpdateTag("aaaaaaaaa")
                .withPublisher("{extension-Publisher}").withTypePropertiesType("{extension-Type}")
                .withTypeHandlerVersion("{handler-version}").withAutoUpgradeMinorVersion(true)
                .withEnableAutomaticUpgrade(true)
                .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize("{}",
                    Object.class, SerializerEncoding.JSON))
                .withProtectedSettings(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize("{}",
                    Object.class, SerializerEncoding.JSON))
                .withProvisionAfterExtensions(Arrays.asList("aa")).withSuppressFailures(true),
            com.azure.core.util.Context.NONE);
    }
}
