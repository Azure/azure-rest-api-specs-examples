
import com.azure.resourcemanager.containerinstance.fluent.models.ContainerGroupProfileInner;
import com.azure.resourcemanager.containerinstance.models.Container;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupIpAddressType;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupNetworkProtocol;
import com.azure.resourcemanager.containerinstance.models.ContainerPort;
import com.azure.resourcemanager.containerinstance.models.EncryptionProperties;
import com.azure.resourcemanager.containerinstance.models.IpAddress;
import com.azure.resourcemanager.containerinstance.models.OperatingSystemTypes;
import com.azure.resourcemanager.containerinstance.models.Port;
import com.azure.resourcemanager.containerinstance.models.ResourceRequests;
import com.azure.resourcemanager.containerinstance.models.ResourceRequirements;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CGProfile CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerGroupProfileCreateOrUpdate_EncryptionProperties.json
     */
    /**
     * Sample code: ContainerGroupProfileWithEncryptionProperties.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerGroupProfileWithEncryptionProperties(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getCGProfiles().createOrUpdateWithResponse("demo", "demo1",
            new ContainerGroupProfileInner().withLocation("eastus2").withZones(Arrays.asList("1"))
                .withEncryptionProperties(new EncryptionProperties()
                    .withVaultBaseUrl("https://testkeyvault.vault.azure.net").withKeyName("fakeTokenPlaceholder")
                    .withKeyVersion("fakeTokenPlaceholder").withIdentity(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity"))
                .withContainers(Arrays.asList(new Container().withName("demo1").withImage("nginx")
                    .withCommand(Arrays.asList()).withPorts(Arrays.asList(new ContainerPort().withPort(80)))
                    .withEnvironmentVariables(Arrays.asList())
                    .withResources(new ResourceRequirements()
                        .withRequests(new ResourceRequests().withMemoryInGB(1.5).withCpu(1.0)))))
                .withImageRegistryCredentials(Arrays.asList())
                .withIpAddress(new IpAddress()
                    .withPorts(Arrays.asList(new Port().withProtocol(ContainerGroupNetworkProtocol.TCP).withPort(80)))
                    .withType(ContainerGroupIpAddressType.PUBLIC))
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
