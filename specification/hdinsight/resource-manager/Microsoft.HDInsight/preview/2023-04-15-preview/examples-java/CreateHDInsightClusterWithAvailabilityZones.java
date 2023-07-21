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
import com.azure.resourcemanager.hdinsight.models.SshProfile;
import com.azure.resourcemanager.hdinsight.models.SshPublicKey;
import com.azure.resourcemanager.hdinsight.models.StorageAccount;
import com.azure.resourcemanager.hdinsight.models.StorageProfile;
import com.azure.resourcemanager.hdinsight.models.VirtualNetworkProfile;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Clusters Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/CreateHDInsightClusterWithAvailabilityZones.json
     */
    /**
     * Sample code: Create cluster with availability zones.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void createClusterWithAvailabilityZones(com.azure.resourcemanager.hdinsight.HDInsightManager manager)
        throws IOException {
        manager
            .clusters()
            .define("cluster1")
            .withExistingResourceGroup("rg1")
            .withZones(Arrays.asList("1"))
            .withProperties(
                new ClusterCreateProperties()
                    .withClusterVersion("3.6")
                    .withOsType(OSType.LINUX)
                    .withClusterDefinition(
                        new ClusterDefinition()
                            .withKind("hadoop")
                            .withConfigurations(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize(
                                        "{\"ambari-conf\":{\"database-name\":\"{ambari database"
                                            + " name}\",\"database-server\":\"{sql server"
                                            + " name}.database.windows.net\",\"database-user-name\":\"**********\",\"database-user-password\":\"**********\"},\"gateway\":{\"restAuthCredential.isEnabled\":true,\"restAuthCredential.password\":\"**********\",\"restAuthCredential.username\":\"admin\"},\"hive-env\":{\"hive_database\":\"Existing"
                                            + " MSSQL Server database with SQL"
                                            + " authentication\",\"hive_database_name\":\"{hive metastore"
                                            + " name}\",\"hive_database_type\":\"mssql\",\"hive_existing_mssql_server_database\":\"{hive"
                                            + " metastore name}\",\"hive_existing_mssql_server_host\":\"{sql server"
                                            + " name}.database.windows.net\",\"hive_hostname\":\"{sql server"
                                            + " name}.database.windows.net\"},\"hive-site\":{\"javax.jdo.option.ConnectionDriverName\":\"com.microsoft.sqlserver.jdbc.SQLServerDriver\",\"javax.jdo.option.ConnectionPassword\":\"**********!\",\"javax.jdo.option.ConnectionURL\":\"jdbc:sqlserver://{sql"
                                            + " server name}.database.windows.net;database={hive metastore"
                                            + " name};encrypt=true;trustServerCertificate=true;create=false;loginTimeout=300;sendStringParametersAsUnicode=true;prepareSQL=0\",\"javax.jdo.option.ConnectionUserName\":\"**********\"},\"oozie-env\":{\"oozie_database\":\"Existing"
                                            + " MSSQL Server database with SQL"
                                            + " authentication\",\"oozie_database_name\":\"{oozie metastore"
                                            + " name}\",\"oozie_database_type\":\"mssql\",\"oozie_existing_mssql_server_database\":\"{oozie"
                                            + " metastore name}\",\"oozie_existing_mssql_server_host\":\"{sql server"
                                            + " name}.database.windows.net\",\"oozie_hostname\":\"{sql server"
                                            + " name}.database.windows.net\"},\"oozie-site\":{\"oozie.db.schema.name\":\"oozie\",\"oozie.service.JPAService.jdbc.driver\":\"com.microsoft.sqlserver.jdbc.SQLServerDriver\",\"oozie.service.JPAService.jdbc.password\":\"**********\",\"oozie.service.JPAService.jdbc.url\":\"jdbc:sqlserver://{sql"
                                            + " server name}.database.windows.net;database={oozie metastore"
                                            + " name};encrypt=true;trustServerCertificate=true;create=false;loginTimeout=300;sendStringParametersAsUnicode=true;prepareSQL=0\",\"oozie.service.JPAService.jdbc.username\":\"**********\"}}",
                                        Object.class,
                                        SerializerEncoding.JSON)))
                    .withComputeProfile(
                        new ComputeProfile()
                            .withRoles(
                                Arrays
                                    .asList(
                                        new Role()
                                            .withName("headnode")
                                            .withTargetInstanceCount(2)
                                            .withHardwareProfile(new HardwareProfile().withVmSize("standard_d3"))
                                            .withOsProfile(
                                                new OsProfile()
                                                    .withLinuxOperatingSystemProfile(
                                                        new LinuxOperatingSystemProfile()
                                                            .withUsername("sshuser")
                                                            .withPassword("fakeTokenPlaceholder")
                                                            .withSshProfile(
                                                                new SshProfile()
                                                                    .withPublicKeys(
                                                                        Arrays
                                                                            .asList(
                                                                                new SshPublicKey()
                                                                                    .withCertificateData(
                                                                                        "**********"))))))
                                            .withVirtualNetworkProfile(
                                                new VirtualNetworkProfile()
                                                    .withId(
                                                        "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname")
                                                    .withSubnet(
                                                        "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet")),
                                        new Role()
                                            .withName("workernode")
                                            .withTargetInstanceCount(2)
                                            .withHardwareProfile(new HardwareProfile().withVmSize("standard_d3"))
                                            .withOsProfile(
                                                new OsProfile()
                                                    .withLinuxOperatingSystemProfile(
                                                        new LinuxOperatingSystemProfile()
                                                            .withUsername("sshuser")
                                                            .withPassword("fakeTokenPlaceholder")
                                                            .withSshProfile(
                                                                new SshProfile()
                                                                    .withPublicKeys(
                                                                        Arrays
                                                                            .asList(
                                                                                new SshPublicKey()
                                                                                    .withCertificateData(
                                                                                        "**********"))))))
                                            .withVirtualNetworkProfile(
                                                new VirtualNetworkProfile()
                                                    .withId(
                                                        "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname")
                                                    .withSubnet(
                                                        "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet")))))
                    .withStorageProfile(
                        new StorageProfile()
                            .withStorageaccounts(
                                Arrays
                                    .asList(
                                        new StorageAccount()
                                            .withName("mystorage")
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
