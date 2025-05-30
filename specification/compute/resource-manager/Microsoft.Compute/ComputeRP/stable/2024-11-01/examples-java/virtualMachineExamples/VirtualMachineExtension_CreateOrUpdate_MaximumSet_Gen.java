
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineExtensionInner;
import com.azure.resourcemanager.compute.models.InstanceViewStatus;
import com.azure.resourcemanager.compute.models.StatusLevelTypes;
import com.azure.resourcemanager.compute.models.VirtualMachineExtensionInstanceView;
import java.io.IOException;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachineExtensions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineExamples/VirtualMachineExtension_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtension_CreateOrUpdate_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionCreateOrUpdateMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) throws IOException {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineExtensions().createOrUpdate("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaa",
            new VirtualMachineExtensionInner().withLocation("westus").withTags(mapOf("key9183", "fakeTokenPlaceholder"))
                .withForceUpdateTag("a").withPublisher("extPublisher").withTypePropertiesType("extType")
                .withTypeHandlerVersion("1.2").withAutoUpgradeMinorVersion(true).withEnableAutomaticUpgrade(true)
                .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize("{}",
                    Object.class, SerializerEncoding.JSON))
                .withProtectedSettings(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize("{}",
                    Object.class, SerializerEncoding.JSON))
                .withInstanceView(new VirtualMachineExtensionInstanceView().withName("aaaaaaaaaaaaaaaaa")
                    .withType("aaaaaaaaa").withTypeHandlerVersion("aaaaaaaaaaaaaaaaaaaaaaaaaa")
                    .withSubstatuses(Arrays.asList(new InstanceViewStatus().withCode("fakeTokenPlaceholder")
                        .withLevel(StatusLevelTypes.INFO).withDisplayStatus("aaaaaa").withMessage("a")
                        .withTime(OffsetDateTime.parse("2021-11-30T12:58:26.522Z"))))
                    .withStatuses(Arrays.asList(new InstanceViewStatus().withCode("fakeTokenPlaceholder")
                        .withLevel(StatusLevelTypes.INFO).withDisplayStatus("aaaaaa").withMessage("a")
                        .withTime(OffsetDateTime.parse("2021-11-30T12:58:26.522Z")))))
                .withSuppressFailures(true),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
