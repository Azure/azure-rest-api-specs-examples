
import com.azure.resourcemanager.containerinstance.fluent.models.ContainerGroupInner;
import com.azure.resourcemanager.containerinstance.models.ConfidentialComputeProperties;
import com.azure.resourcemanager.containerinstance.models.Container;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupIpAddressType;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupNetworkProtocol;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupSku;
import com.azure.resourcemanager.containerinstance.models.ContainerPort;
import com.azure.resourcemanager.containerinstance.models.IpAddress;
import com.azure.resourcemanager.containerinstance.models.OperatingSystemTypes;
import com.azure.resourcemanager.containerinstance.models.Port;
import com.azure.resourcemanager.containerinstance.models.ResourceRequests;
import com.azure.resourcemanager.containerinstance.models.ResourceRequirements;
import com.azure.resourcemanager.containerinstance.models.SecurityContextCapabilitiesDefinition;
import com.azure.resourcemanager.containerinstance.models.SecurityContextDefinition;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ContainerGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/
     * ContainerGroupCreateConfidential.json
     */
    /**
     * Sample code: ConfidentialContainerGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void confidentialContainerGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainerGroups().createOrUpdate("demo", "demo1",
            new ContainerGroupInner().withLocation("westeurope")
                .withContainers(Arrays.asList(new Container().withName("accdemo").withImage("confiimage")
                    .withCommand(Arrays.asList()).withPorts(Arrays.asList(new ContainerPort().withPort(8000)))
                    .withEnvironmentVariables(Arrays.asList())
                    .withResources(new ResourceRequirements()
                        .withRequests(new ResourceRequests().withMemoryInGB(1.5).withCpu(1.0)))
                    .withSecurityContext(new SecurityContextDefinition().withPrivileged(false).withCapabilities(
                        new SecurityContextCapabilitiesDefinition().withAdd(Arrays.asList("CAP_NET_ADMIN"))))))
                .withImageRegistryCredentials(Arrays.asList())
                .withIpAddress(new IpAddress()
                    .withPorts(Arrays.asList(new Port().withProtocol(ContainerGroupNetworkProtocol.TCP).withPort(8000)))
                    .withType(ContainerGroupIpAddressType.PUBLIC))
                .withOsType(OperatingSystemTypes.LINUX).withSku(ContainerGroupSku.CONFIDENTIAL)
                .withConfidentialComputeProperties(new ConfidentialComputeProperties().withCcePolicy(
                    "eyJhbGxvd19hbGwiOiB0cnVlLCAiY29udGFpbmVycyI6IHsibGVuZ3RoIjogMCwgImVsZW1lbnRzIjogbnVsbH19")),
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
