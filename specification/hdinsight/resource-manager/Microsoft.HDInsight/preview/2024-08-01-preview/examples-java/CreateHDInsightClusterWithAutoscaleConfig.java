
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.hdinsight.models.Autoscale;
import com.azure.resourcemanager.hdinsight.models.AutoscaleRecurrence;
import com.azure.resourcemanager.hdinsight.models.AutoscaleSchedule;
import com.azure.resourcemanager.hdinsight.models.AutoscaleTimeAndCapacity;
import com.azure.resourcemanager.hdinsight.models.ClusterCreateProperties;
import com.azure.resourcemanager.hdinsight.models.ClusterDefinition;
import com.azure.resourcemanager.hdinsight.models.ComputeProfile;
import com.azure.resourcemanager.hdinsight.models.DaysOfWeek;
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

/**
 * Samples for Clusters Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * CreateHDInsightClusterWithAutoscaleConfig.json
     */
    /**
     * Sample code: Create HDInsight cluster with Autoscale configuration.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void createHDInsightClusterWithAutoscaleConfiguration(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) throws IOException {
        manager.clusters().define("cluster1").withExistingResourceGroup("rg1")
            .withProperties(new ClusterCreateProperties().withClusterVersion("3.6").withOsType(OSType.LINUX)
                .withTier(Tier.STANDARD)
                .withClusterDefinition(new ClusterDefinition().withKind("hadoop")
                    .withComponentVersion(mapOf("Hadoop", "2.7"))
                    .withConfigurations(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                        "{\"gateway\":{\"restAuthCredential.isEnabled\":true,\"restAuthCredential.password\":\"**********\",\"restAuthCredential.username\":\"admin\"}}",
                        Object.class, SerializerEncoding.JSON)))
                .withComputeProfile(new ComputeProfile()
                    .withRoles(Arrays.asList(new Role().withName("workernode").withTargetInstanceCount(4)
                        .withAutoscaleConfiguration(new Autoscale().withRecurrence(new AutoscaleRecurrence()
                            .withTimeZone("China Standard Time")
                            .withSchedule(Arrays.asList(
                                new AutoscaleSchedule()
                                    .withDays(Arrays.asList(DaysOfWeek.MONDAY, DaysOfWeek.TUESDAY, DaysOfWeek.WEDNESDAY,
                                        DaysOfWeek.THURSDAY, DaysOfWeek.FRIDAY))
                                    .withTimeAndCapacity(new AutoscaleTimeAndCapacity().withTime("09:00")
                                        .withMinInstanceCount(3).withMaxInstanceCount(3)),
                                new AutoscaleSchedule()
                                    .withDays(Arrays.asList(DaysOfWeek.MONDAY, DaysOfWeek.TUESDAY, DaysOfWeek.WEDNESDAY,
                                        DaysOfWeek.THURSDAY, DaysOfWeek.FRIDAY))
                                    .withTimeAndCapacity(new AutoscaleTimeAndCapacity().withTime("18:00")
                                        .withMinInstanceCount(6).withMaxInstanceCount(6)),
                                new AutoscaleSchedule().withDays(Arrays.asList(DaysOfWeek.SATURDAY, DaysOfWeek.SUNDAY))
                                    .withTimeAndCapacity(new AutoscaleTimeAndCapacity()
                                        .withTime("09:00").withMinInstanceCount(2).withMaxInstanceCount(2)),
                                new AutoscaleSchedule().withDays(Arrays.asList(DaysOfWeek.SATURDAY, DaysOfWeek.SUNDAY))
                                    .withTimeAndCapacity(new AutoscaleTimeAndCapacity().withTime("18:00")
                                        .withMinInstanceCount(4).withMaxInstanceCount(4))))))
                        .withHardwareProfile(new HardwareProfile().withVmSize("Standard_D4_V2"))
                        .withOsProfile(new OsProfile().withLinuxOperatingSystemProfile(new LinuxOperatingSystemProfile()
                            .withUsername("sshuser").withPassword("fakeTokenPlaceholder")))
                        .withScriptActions(Arrays.asList()))))
                .withStorageProfile(new StorageProfile()
                    .withStorageaccounts(Arrays.asList(new StorageAccount().withName("mystorage.blob.core.windows.net")
                        .withIsDefault(true).withContainer("hdinsight-autoscale-tes-2019-06-18t05-49-16-591z")
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
