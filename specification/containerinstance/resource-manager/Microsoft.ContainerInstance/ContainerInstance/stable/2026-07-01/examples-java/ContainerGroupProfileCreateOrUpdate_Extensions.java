
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.containerinstance.fluent.models.ContainerGroupProfileInner;
import com.azure.resourcemanager.containerinstance.models.Container;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupIpAddressType;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupNetworkProtocol;
import com.azure.resourcemanager.containerinstance.models.ContainerPort;
import com.azure.resourcemanager.containerinstance.models.DeploymentExtensionSpec;
import com.azure.resourcemanager.containerinstance.models.IpAddress;
import com.azure.resourcemanager.containerinstance.models.OperatingSystemTypes;
import com.azure.resourcemanager.containerinstance.models.Port;
import com.azure.resourcemanager.containerinstance.models.ResourceRequests;
import com.azure.resourcemanager.containerinstance.models.ResourceRequirements;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CGProfile CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupProfileCreateOrUpdate_Extensions.json
     */
    /**
     * Sample code: ContainerGroupProfileCreateWithExtensions.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerGroupProfileCreateWithExtensions(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) throws IOException {
        manager.serviceClient().getCGProfiles().createOrUpdateWithResponse("demo", "demo1",
            new ContainerGroupProfileInner().withLocation("eastus2").withZones(Arrays.asList("1"))
                .withContainers(Arrays.asList(new Container().withName("demo1").withImage("nginx")
                    .withCommand(Arrays.asList()).withPorts(Arrays.asList(new ContainerPort().withPort(80)))
                    .withEnvironmentVariables(Arrays.asList())
                    .withResources(new ResourceRequirements()
                        .withRequests(new ResourceRequests().withMemoryInGB(1.5).withCpu(1.0)))))
                .withExtensions(Arrays.asList(
                    new DeploymentExtensionSpec().withName("kube-proxy").withExtensionType("kube-proxy")
                        .withVersion("1.0")
                        .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                            "{\"clusterCidr\":\"10.240.0.0/16\",\"kubeVersion\":\"v1.9.10\"}", Object.class,
                            SerializerEncoding.JSON))
                        .withProtectedSettings(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                            "{\"kubeConfig\":\"<kubeconfig encoded string>\"}", Object.class, SerializerEncoding.JSON)),
                    new DeploymentExtensionSpec()
                        .withName("vk-realtime-metrics").withExtensionType("realtime-metrics").withVersion("1.0")))
                .withImageRegistryCredentials(Arrays.asList())
                .withIpAddress(new IpAddress()
                    .withPorts(Arrays.asList(new Port().withProtocol(ContainerGroupNetworkProtocol.TCP).withPort(80)))
                    .withType(ContainerGroupIpAddressType.PRIVATE))
                .withOsType(OperatingSystemTypes.LINUX),
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
