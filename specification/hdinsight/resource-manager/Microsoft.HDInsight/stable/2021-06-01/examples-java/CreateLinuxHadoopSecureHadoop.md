Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.hdinsight.models.ClusterCreateProperties;
import com.azure.resourcemanager.hdinsight.models.ClusterDefinition;
import com.azure.resourcemanager.hdinsight.models.ComputeProfile;
import com.azure.resourcemanager.hdinsight.models.DirectoryType;
import com.azure.resourcemanager.hdinsight.models.HardwareProfile;
import com.azure.resourcemanager.hdinsight.models.LinuxOperatingSystemProfile;
import com.azure.resourcemanager.hdinsight.models.OSType;
import com.azure.resourcemanager.hdinsight.models.OsProfile;
import com.azure.resourcemanager.hdinsight.models.Role;
import com.azure.resourcemanager.hdinsight.models.SecurityProfile;
import com.azure.resourcemanager.hdinsight.models.SshProfile;
import com.azure.resourcemanager.hdinsight.models.SshPublicKey;
import com.azure.resourcemanager.hdinsight.models.StorageAccount;
import com.azure.resourcemanager.hdinsight.models.StorageProfile;
import com.azure.resourcemanager.hdinsight.models.Tier;
import com.azure.resourcemanager.hdinsight.models.VirtualNetworkProfile;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Clusters Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/CreateLinuxHadoopSecureHadoop.json
     */
    /**
     * Sample code: Create Secure Hadoop cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void createSecureHadoopCluster(com.azure.resourcemanager.hdinsight.HDInsightManager manager)
        throws IOException {
        manager
            .clusters()
            .define("cluster1")
            .withExistingResourceGroup("rg1")
            .withTags(mapOf("key1", "val1"))
            .withProperties(
                new ClusterCreateProperties()
                    .withClusterVersion("3.5")
                    .withOsType(OSType.LINUX)
                    .withTier(Tier.PREMIUM)
                    .withClusterDefinition(
                        new ClusterDefinition()
                            .withKind("Hadoop")
                            .withConfigurations(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize(
                                        "{\"gateway\":{\"restAuthCredential.isEnabled\":true,\"restAuthCredential.password\":\"**********\",\"restAuthCredential.username\":\"admin\"}}",
                                        Object.class,
                                        SerializerEncoding.JSON)))
                    .withSecurityProfile(
                        new SecurityProfile()
                            .withDirectoryType(DirectoryType.ACTIVE_DIRECTORY)
                            .withDomain("DomainName")
                            .withOrganizationalUnitDN("OU=Hadoop,DC=hdinsight,DC=test")
                            .withLdapsUrls(Arrays.asList("ldaps://10.10.0.4:636"))
                            .withDomainUsername("DomainUsername")
                            .withDomainUserPassword("**********")
                            .withClusterUsersGroupDNs(Arrays.asList("hdiusers")))
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
                                                            .withPassword("**********")
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
                                                        "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"))
                                            .withScriptActions(Arrays.asList()),
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
                                                            .withPassword("**********")
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
                                                        "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"))
                                            .withScriptActions(Arrays.asList()),
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
                                                            .withPassword("**********")
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
                                                        "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet"))
                                            .withScriptActions(Arrays.asList()))))
                    .withStorageProfile(
                        new StorageProfile()
                            .withStorageaccounts(
                                Arrays
                                    .asList(
                                        new StorageAccount()
                                            .withName("mystorage.blob.core.windows.net")
                                            .withIsDefault(true)
                                            .withContainer("containername")
                                            .withKey("storage account key")))))
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
```
