
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.hdinsight.models.ClientGroupInfo;
import com.azure.resourcemanager.hdinsight.models.ClusterCreateProperties;
import com.azure.resourcemanager.hdinsight.models.ClusterDefinition;
import com.azure.resourcemanager.hdinsight.models.ComputeProfile;
import com.azure.resourcemanager.hdinsight.models.DataDisksGroups;
import com.azure.resourcemanager.hdinsight.models.HardwareProfile;
import com.azure.resourcemanager.hdinsight.models.KafkaRestProperties;
import com.azure.resourcemanager.hdinsight.models.LinuxOperatingSystemProfile;
import com.azure.resourcemanager.hdinsight.models.OSType;
import com.azure.resourcemanager.hdinsight.models.OsProfile;
import com.azure.resourcemanager.hdinsight.models.Role;
import com.azure.resourcemanager.hdinsight.models.StorageAccount;
import com.azure.resourcemanager.hdinsight.models.StorageProfile;
import com.azure.resourcemanager.hdinsight.models.Tier;
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
     * CreateKafkaClusterWithKafkaRestProxy.json
     */
    /**
     * Sample code: Create Kafka cluster with Kafka Rest Proxy.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void createKafkaClusterWithKafkaRestProxy(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) throws IOException {
        manager.clusters().define("cluster1").withExistingResourceGroup("rg1")
            .withProperties(new ClusterCreateProperties().withClusterVersion("4.0").withOsType(OSType.LINUX)
                .withTier(Tier.STANDARD)
                .withClusterDefinition(new ClusterDefinition().withKind("kafka")
                    .withComponentVersion(mapOf("Kafka", "2.1"))
                    .withConfigurations(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                        "{\"gateway\":{\"restAuthCredential.isEnabled\":true,\"restAuthCredential.password\":\"**********\",\"restAuthCredential.username\":\"admin\"}}",
                        Object.class, SerializerEncoding.JSON)))
                .withKafkaRestProperties(new KafkaRestProperties().withClientGroupInfo(new ClientGroupInfo()
                    .withGroupName("Kafka security group name").withGroupId("00000000-0000-0000-0000-111111111111")))
                .withComputeProfile(new ComputeProfile().withRoles(Arrays.asList(
                    new Role().withName("headnode").withTargetInstanceCount(2)
                        .withHardwareProfile(new HardwareProfile().withVmSize("Large"))
                        .withOsProfile(
                            new OsProfile().withLinuxOperatingSystemProfile(new LinuxOperatingSystemProfile()
                                .withUsername("sshuser").withPassword("fakeTokenPlaceholder"))),
                    new Role().withName("workernode").withTargetInstanceCount(3)
                        .withHardwareProfile(new HardwareProfile().withVmSize("Large"))
                        .withOsProfile(new OsProfile().withLinuxOperatingSystemProfile(new LinuxOperatingSystemProfile()
                            .withUsername("sshuser").withPassword("fakeTokenPlaceholder")))
                        .withDataDisksGroups(Arrays.asList(new DataDisksGroups().withDisksPerNode(8))),
                    new Role().withName("zookeepernode").withTargetInstanceCount(3)
                        .withHardwareProfile(new HardwareProfile().withVmSize("Small"))
                        .withOsProfile(new OsProfile().withLinuxOperatingSystemProfile(new LinuxOperatingSystemProfile()
                            .withUsername("sshuser").withPassword("fakeTokenPlaceholder"))),
                    new Role().withName("kafkamanagementnode").withTargetInstanceCount(2)
                        .withHardwareProfile(new HardwareProfile().withVmSize("Standard_D4_v2"))
                        .withOsProfile(new OsProfile().withLinuxOperatingSystemProfile(new LinuxOperatingSystemProfile()
                            .withUsername("kafkauser").withPassword("fakeTokenPlaceholder"))))))
                .withStorageProfile(new StorageProfile().withStorageaccounts(Arrays.asList(new StorageAccount()
                    .withName("mystorage.blob.core.windows.net").withIsDefault(true).withContainer("containername")
                    .withKey("fakeTokenPlaceholder").withEnableSecureChannel(true)))))
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
