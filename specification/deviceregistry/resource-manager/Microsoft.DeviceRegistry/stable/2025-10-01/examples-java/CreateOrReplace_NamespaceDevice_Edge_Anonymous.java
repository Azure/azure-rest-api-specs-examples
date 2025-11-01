
import com.azure.core.util.BinaryData;
import com.azure.resourcemanager.deviceregistry.models.AuthenticationMethod;
import com.azure.resourcemanager.deviceregistry.models.ExtendedLocation;
import com.azure.resourcemanager.deviceregistry.models.HostAuthentication;
import com.azure.resourcemanager.deviceregistry.models.InboundEndpoints;
import com.azure.resourcemanager.deviceregistry.models.MessagingEndpoints;
import com.azure.resourcemanager.deviceregistry.models.NamespaceDeviceProperties;
import java.nio.charset.StandardCharsets;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NamespaceDevices CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/CreateOrReplace_NamespaceDevice_Edge_Anonymous.json
     */
    /**
     * Sample code: Create edge enabled device with anonymous host authentication.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void createEdgeEnabledDeviceWithAnonymousHostAuthentication(
        com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDevices().define("namespace-device-on-edge").withRegion("West Europe")
            .withExistingNamespace("myResourceGroup", "adr-namespace-gbk0925-n01")
            .withProperties(
                new NamespaceDeviceProperties().withEnabled(true).withExternalDeviceId("unique-edge-device-identifier")
                    .withEndpoints(new MessagingEndpoints().withInbound(mapOf("theOnlyOPCUABroker",
                        new InboundEndpoints().withEndpointType("microsoft.opcua")
                            .withAddress("opc.tcp://192.168.86.23:51211/UA/SampleServer").withVersion("2")
                            .withAuthentication(new HostAuthentication().withMethod(AuthenticationMethod.ANONYMOUS)))))
                    .withAttributes(
                        mapOf("deviceType", BinaryData.fromBytes("dough-maker".getBytes(StandardCharsets.UTF_8)),
                            "deviceOwner", BinaryData.fromBytes("OT".getBytes(StandardCharsets.UTF_8)),
                            "deviceCategory", BinaryData.fromBytes("16".getBytes(StandardCharsets.UTF_8)))))
            .withExtendedLocation(new ExtendedLocation().withType("CustomLocation").withName(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"))
            .create();
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
