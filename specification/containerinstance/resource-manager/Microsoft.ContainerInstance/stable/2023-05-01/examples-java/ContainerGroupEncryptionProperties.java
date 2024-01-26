
import com.azure.resourcemanager.containerinstance.fluent.models.ContainerGroupInner;
import com.azure.resourcemanager.containerinstance.models.Container;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupIdentity;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupIdentityUserAssignedIdentities;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupIpAddressType;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupNetworkProtocol;
import com.azure.resourcemanager.containerinstance.models.ContainerPort;
import com.azure.resourcemanager.containerinstance.models.EncryptionProperties;
import com.azure.resourcemanager.containerinstance.models.IpAddress;
import com.azure.resourcemanager.containerinstance.models.OperatingSystemTypes;
import com.azure.resourcemanager.containerinstance.models.Port;
import com.azure.resourcemanager.containerinstance.models.ResourceIdentityType;
import com.azure.resourcemanager.containerinstance.models.ResourceRequests;
import com.azure.resourcemanager.containerinstance.models.ResourceRequirements;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ContainerGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/
     * ContainerGroupEncryptionProperties.json
     */
    /**
     * Sample code: ContainerGroupWithEncryptionProperties.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupWithEncryptionProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainerGroups().createOrUpdate("demo", "demo1",
            new ContainerGroupInner().withLocation("eastus2").withIdentity(new ContainerGroupIdentity()
                .withType(ResourceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity",
                    new ContainerGroupIdentityUserAssignedIdentities())))
                .withContainers(Arrays.asList(new Container().withName("demo1").withImage("nginx")
                    .withCommand(Arrays.asList()).withPorts(Arrays.asList(new ContainerPort().withPort(80)))
                    .withEnvironmentVariables(Arrays.asList())
                    .withResources(new ResourceRequirements()
                        .withRequests(new ResourceRequests().withMemoryInGB(1.5).withCpu(1.0)))))
                .withImageRegistryCredentials(Arrays.asList())
                .withIpAddress(new IpAddress()
                    .withPorts(Arrays.asList(new Port().withProtocol(ContainerGroupNetworkProtocol.TCP).withPort(80)))
                    .withType(ContainerGroupIpAddressType.PUBLIC))
                .withOsType(OperatingSystemTypes.LINUX)
                .withEncryptionProperties(new EncryptionProperties()
                    .withVaultBaseUrl("https://testkeyvault.vault.azure.net").withKeyName("fakeTokenPlaceholder")
                    .withKeyVersion("fakeTokenPlaceholder").withIdentity(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity")),
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
