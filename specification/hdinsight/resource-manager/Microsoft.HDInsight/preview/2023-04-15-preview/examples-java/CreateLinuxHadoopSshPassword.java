import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.hdinsight.models.ClusterCreateProperties;
import com.azure.resourcemanager.hdinsight.models.ClusterDefinition;
import com.azure.resourcemanager.hdinsight.models.ComputeProfile;
import com.azure.resourcemanager.hdinsight.models.HardwareProfile;
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

/** Samples for Clusters Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/CreateLinuxHadoopSshPassword.json
     */
    /**
     * Sample code: Create Hadoop on Linux cluster with SSH password.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void createHadoopOnLinuxClusterWithSSHPassword(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) throws IOException {
        manager
            .clusters()
            .define("cluster1")
            .withExistingResourceGroup("rg1")
            .withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withProperties(
                new ClusterCreateProperties()
                    .withClusterVersion("3.5")
                    .withOsType(OSType.LINUX)
                    .withTier(Tier.STANDARD)
                    .withClusterDefinition(
                        new ClusterDefinition()
                            .withKind("Hadoop")
                            .withConfigurations(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize(
                                        "{\"gateway\":{\"restAuthCredential.isEnabled\":\"true\",\"restAuthCredential.password\":\"**********\",\"restAuthCredential.username\":\"admin\"}}",
                                        Object.class,
                                        SerializerEncoding.JSON)))
                    .withComputeProfile(
                        new ComputeProfile()
                            .withRoles(
                                Arrays
                                    .asList(
                                        new Role()
                                            .withName("headnode")
                                            .withMinInstanceCount(1)
                                            .withTargetInstanceCount(2)
                                            .withHardwareProfile(new HardwareProfile().withVmSize("Standard_D3_V2"))
                                            .withOsProfile(
                                                new OsProfile()
                                                    .withLinuxOperatingSystemProfile(
                                                        new LinuxOperatingSystemProfile()
                                                            .withUsername("sshuser")
                                                            .withPassword("fakeTokenPlaceholder"))),
                                        new Role()
                                            .withName("workernode")
                                            .withMinInstanceCount(1)
                                            .withTargetInstanceCount(4)
                                            .withHardwareProfile(new HardwareProfile().withVmSize("Standard_D3_V2"))
                                            .withOsProfile(
                                                new OsProfile()
                                                    .withLinuxOperatingSystemProfile(
                                                        new LinuxOperatingSystemProfile()
                                                            .withUsername("sshuser")
                                                            .withPassword("fakeTokenPlaceholder"))),
                                        new Role()
                                            .withName("zookeepernode")
                                            .withMinInstanceCount(1)
                                            .withTargetInstanceCount(3)
                                            .withHardwareProfile(new HardwareProfile().withVmSize("Small"))
                                            .withOsProfile(
                                                new OsProfile()
                                                    .withLinuxOperatingSystemProfile(
                                                        new LinuxOperatingSystemProfile()
                                                            .withUsername("sshuser")
                                                            .withPassword("fakeTokenPlaceholder"))))))
                    .withStorageProfile(
                        new StorageProfile()
                            .withStorageaccounts(
                                Arrays
                                    .asList(
                                        new StorageAccount()
                                            .withName("mystorage.blob.core.windows.net")
                                            .withIsDefault(true)
                                            .withContainer("containername")
                                            .withKey("fakeTokenPlaceholder")
                                            .withEnableSecureChannel(true)))))
            .create();
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
