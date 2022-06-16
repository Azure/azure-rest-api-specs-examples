import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.containerinstance.fluent.models.ContainerGroupInner;
import com.azure.resourcemanager.containerinstance.models.AzureFileVolume;
import com.azure.resourcemanager.containerinstance.models.Container;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupDiagnostics;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupIdentity;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupIdentityUserAssignedIdentities;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupIpAddressType;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupNetworkProtocol;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupSubnetId;
import com.azure.resourcemanager.containerinstance.models.ContainerPort;
import com.azure.resourcemanager.containerinstance.models.DnsConfiguration;
import com.azure.resourcemanager.containerinstance.models.GpuResource;
import com.azure.resourcemanager.containerinstance.models.GpuSku;
import com.azure.resourcemanager.containerinstance.models.IpAddress;
import com.azure.resourcemanager.containerinstance.models.LogAnalytics;
import com.azure.resourcemanager.containerinstance.models.LogAnalyticsLogType;
import com.azure.resourcemanager.containerinstance.models.OperatingSystemTypes;
import com.azure.resourcemanager.containerinstance.models.Port;
import com.azure.resourcemanager.containerinstance.models.ResourceIdentityType;
import com.azure.resourcemanager.containerinstance.models.ResourceRequests;
import com.azure.resourcemanager.containerinstance.models.ResourceRequirements;
import com.azure.resourcemanager.containerinstance.models.Volume;
import com.azure.resourcemanager.containerinstance.models.VolumeMount;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ContainerGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerGroupsCreateOrUpdate.json
     */
    /**
     * Sample code: ContainerGroupsCreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupsCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure
            .containerGroups()
            .manager()
            .serviceClient()
            .getContainerGroups()
            .createOrUpdate(
                "demo",
                "demo1",
                new ContainerGroupInner()
                    .withLocation("west us")
                    .withIdentity(
                        new ContainerGroupIdentity()
                            .withType(ResourceIdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                            .withUserAssignedIdentities(
                                mapOf(
                                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-name",
                                    new ContainerGroupIdentityUserAssignedIdentities())))
                    .withContainers(
                        Arrays
                            .asList(
                                new Container()
                                    .withName("demo1")
                                    .withImage("nginx")
                                    .withCommand(Arrays.asList())
                                    .withPorts(Arrays.asList(new ContainerPort().withPort(80)))
                                    .withEnvironmentVariables(Arrays.asList())
                                    .withResources(
                                        new ResourceRequirements()
                                            .withRequests(
                                                new ResourceRequests()
                                                    .withMemoryInGB(1.5)
                                                    .withCpu(1.0)
                                                    .withGpu(new GpuResource().withCount(1).withSku(GpuSku.K80))))
                                    .withVolumeMounts(
                                        Arrays
                                            .asList(
                                                new VolumeMount()
                                                    .withName("volume1")
                                                    .withMountPath("/mnt/volume1")
                                                    .withReadOnly(false),
                                                new VolumeMount()
                                                    .withName("volume2")
                                                    .withMountPath("/mnt/volume2")
                                                    .withReadOnly(false),
                                                new VolumeMount()
                                                    .withName("volume3")
                                                    .withMountPath("/mnt/volume3")
                                                    .withReadOnly(true)))))
                    .withImageRegistryCredentials(Arrays.asList())
                    .withIpAddress(
                        new IpAddress()
                            .withPorts(
                                Arrays.asList(new Port().withProtocol(ContainerGroupNetworkProtocol.TCP).withPort(80)))
                            .withType(ContainerGroupIpAddressType.PUBLIC)
                            .withDnsNameLabel("dnsnamelabel1"))
                    .withOsType(OperatingSystemTypes.LINUX)
                    .withVolumes(
                        Arrays
                            .asList(
                                new Volume()
                                    .withName("volume1")
                                    .withAzureFile(
                                        new AzureFileVolume()
                                            .withShareName("shareName")
                                            .withStorageAccountName("accountName")
                                            .withStorageAccountKey("accountKey")),
                                new Volume()
                                    .withName("volume2")
                                    .withEmptyDir(
                                        SerializerFactory
                                            .createDefaultManagementSerializerAdapter()
                                            .deserialize("{}", Object.class, SerializerEncoding.JSON)),
                                new Volume()
                                    .withName("volume3")
                                    .withSecret(
                                        mapOf(
                                            "secretKey1",
                                            "SecretValue1InBase64",
                                            "secretKey2",
                                            "SecretValue2InBase64"))))
                    .withDiagnostics(
                        new ContainerGroupDiagnostics()
                            .withLogAnalytics(
                                new LogAnalytics()
                                    .withWorkspaceId("workspaceid")
                                    .withWorkspaceKey("workspaceKey")
                                    .withLogType(LogAnalyticsLogType.CONTAINER_INSIGHTS)
                                    .withMetadata(mapOf("test-key", "test-metadata-value"))
                                    .withWorkspaceResourceId(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/microsoft.operationalinsights/workspaces/workspace")))
                    .withSubnetIds(
                        Arrays
                            .asList(
                                new ContainerGroupSubnetId()
                                    .withId(
                                        "[resourceId('Microsoft.Network/virtualNetworks/subnets',"
                                            + " parameters('vnetName'), parameters('subnetName'))]")))
                    .withDnsConfig(
                        new DnsConfiguration()
                            .withNameServers(Arrays.asList("1.1.1.1"))
                            .withSearchDomains("cluster.local svc.cluster.local")
                            .withOptions("ndots:2")),
                Context.NONE);
    }

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
