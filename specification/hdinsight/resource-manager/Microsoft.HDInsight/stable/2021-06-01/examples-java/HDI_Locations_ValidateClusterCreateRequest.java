import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.hdinsight.models.ClusterCreateProperties;
import com.azure.resourcemanager.hdinsight.models.ClusterCreateRequestValidationParameters;
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

/** Samples for Locations ValidateClusterCreateRequest. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Locations_ValidateClusterCreateRequest.json
     */
    /**
     * Sample code: Get the subscription usages for specific location.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getTheSubscriptionUsagesForSpecificLocation(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) throws IOException {
        manager
            .locations()
            .validateClusterCreateRequestWithResponse(
                "southcentralus",
                new ClusterCreateRequestValidationParameters()
                    .withLocation("southcentralus")
                    .withTags(mapOf())
                    .withProperties(
                        new ClusterCreateProperties()
                            .withClusterVersion("4.0")
                            .withOsType(OSType.LINUX)
                            .withTier(Tier.STANDARD)
                            .withClusterDefinition(
                                new ClusterDefinition()
                                    .withKind("spark")
                                    .withComponentVersion(mapOf("Spark", "2.4"))
                                    .withConfigurations(
                                        SerializerFactory
                                            .createDefaultManagementSerializerAdapter()
                                            .deserialize(
                                                "{\"gateway\":{\"restAuthCredential.isEnabled\":true,\"restAuthCredential.password\":\"**********\",\"restAuthCredential.username\":\"admin\"}}",
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
                                                    .withHardwareProfile(
                                                        new HardwareProfile().withVmSize("Standard_E8_V3"))
                                                    .withOsProfile(
                                                        new OsProfile()
                                                            .withLinuxOperatingSystemProfile(
                                                                new LinuxOperatingSystemProfile()
                                                                    .withUsername("sshuser")
                                                                    .withPassword("********")))
                                                    .withScriptActions(Arrays.asList()),
                                                new Role()
                                                    .withName("workernode")
                                                    .withTargetInstanceCount(4)
                                                    .withHardwareProfile(
                                                        new HardwareProfile().withVmSize("Standard_E8_V3"))
                                                    .withOsProfile(
                                                        new OsProfile()
                                                            .withLinuxOperatingSystemProfile(
                                                                new LinuxOperatingSystemProfile()
                                                                    .withUsername("sshuser")
                                                                    .withPassword("********")))
                                                    .withScriptActions(Arrays.asList()),
                                                new Role()
                                                    .withName("zookeepernode")
                                                    .withMinInstanceCount(1)
                                                    .withTargetInstanceCount(3)
                                                    .withHardwareProfile(
                                                        new HardwareProfile().withVmSize("Standard_D13_V2"))
                                                    .withOsProfile(
                                                        new OsProfile()
                                                            .withLinuxOperatingSystemProfile(
                                                                new LinuxOperatingSystemProfile()
                                                                    .withUsername("sshuser")
                                                                    .withPassword("**********")))
                                                    .withScriptActions(Arrays.asList()))))
                            .withStorageProfile(
                                new StorageProfile()
                                    .withStorageaccounts(
                                        Arrays
                                            .asList(
                                                new StorageAccount()
                                                    .withName("storagename.blob.core.windows.net")
                                                    .withIsDefault(true)
                                                    .withContainer("contianername")
                                                    .withKey("*******")
                                                    .withResourceId(
                                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storagename"))))
                            .withMinSupportedTlsVersion("1.2"))
                    .withName("testclustername")
                    .withType("Microsoft.HDInsight/clusters")
                    .withTenantId("00000000-0000-0000-0000-000000000000")
                    .withFetchAaddsResource(false),
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
