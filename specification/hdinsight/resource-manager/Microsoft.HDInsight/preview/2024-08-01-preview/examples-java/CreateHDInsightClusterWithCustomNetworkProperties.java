
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.hdinsight.models.ClusterCreateProperties;
import com.azure.resourcemanager.hdinsight.models.ClusterDefinition;
import com.azure.resourcemanager.hdinsight.models.ComputeProfile;
import com.azure.resourcemanager.hdinsight.models.HardwareProfile;
import com.azure.resourcemanager.hdinsight.models.IpTag;
import com.azure.resourcemanager.hdinsight.models.LinuxOperatingSystemProfile;
import com.azure.resourcemanager.hdinsight.models.NetworkProperties;
import com.azure.resourcemanager.hdinsight.models.OSType;
import com.azure.resourcemanager.hdinsight.models.OsProfile;
import com.azure.resourcemanager.hdinsight.models.PrivateLink;
import com.azure.resourcemanager.hdinsight.models.ResourceProviderConnection;
import com.azure.resourcemanager.hdinsight.models.Role;
import com.azure.resourcemanager.hdinsight.models.SshProfile;
import com.azure.resourcemanager.hdinsight.models.SshPublicKey;
import com.azure.resourcemanager.hdinsight.models.StorageAccount;
import com.azure.resourcemanager.hdinsight.models.StorageProfile;
import com.azure.resourcemanager.hdinsight.models.VirtualNetworkProfile;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * CreateHDInsightClusterWithCustomNetworkProperties.json
     */
    /**
     * Sample code: Create cluster with network properties.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void createClusterWithNetworkProperties(com.azure.resourcemanager.hdinsight.HDInsightManager manager)
        throws IOException {
        manager.clusters().define("cluster1").withExistingResourceGroup("rg1")
            .withProperties(new ClusterCreateProperties().withClusterVersion("3.6").withOsType(OSType.LINUX)
                .withClusterDefinition(new ClusterDefinition().withKind("hadoop")
                    .withConfigurations(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                        "{\"gateway\":{\"restAuthCredential.isEnabled\":true,\"restAuthCredential.password\":\"**********\",\"restAuthCredential.username\":\"admin\"}}",
                        Object.class, SerializerEncoding.JSON)))
                .withComputeProfile(new ComputeProfile().withRoles(Arrays.asList(
                    new Role().withName("headnode").withTargetInstanceCount(2)
                        .withHardwareProfile(new HardwareProfile().withVmSize("standard_d3"))
                        .withOsProfile(new OsProfile()
                            .withLinuxOperatingSystemProfile(new LinuxOperatingSystemProfile().withUsername("sshuser")
                                .withPassword("fakeTokenPlaceholder")
                                .withSshProfile(new SshProfile()
                                    .withPublicKeys(
                                        Arrays.asList(new SshPublicKey().withCertificateData("**********"))))))
                        .withVirtualNetworkProfile(new VirtualNetworkProfile().withId(
                            "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname")
                            .withSubnet(
                                "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet")),
                    new Role().withName("workernode").withTargetInstanceCount(2)
                        .withHardwareProfile(new HardwareProfile().withVmSize("standard_d3"))
                        .withOsProfile(new OsProfile().withLinuxOperatingSystemProfile(new LinuxOperatingSystemProfile()
                            .withUsername("sshuser").withPassword("fakeTokenPlaceholder")
                            .withSshProfile(new SshProfile()
                                .withPublicKeys(Arrays.asList(new SshPublicKey().withCertificateData("**********"))))))
                        .withVirtualNetworkProfile(new VirtualNetworkProfile().withId(
                            "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname")
                            .withSubnet(
                                "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet")))))
                .withStorageProfile(new StorageProfile()
                    .withStorageaccounts(Arrays.asList(new StorageAccount().withName("mystorage").withIsDefault(true)
                        .withContainer("containername").withKey("fakeTokenPlaceholder").withEnableSecureChannel(true))))
                .withNetworkProperties(
                    new NetworkProperties().withResourceProviderConnection(ResourceProviderConnection.OUTBOUND)
                        .withPrivateLink(PrivateLink.ENABLED)
                        .withPublicIpTag(new IpTag().withIpTagType("FirstPartyUsage").withTag("/<TagName>"))))
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
