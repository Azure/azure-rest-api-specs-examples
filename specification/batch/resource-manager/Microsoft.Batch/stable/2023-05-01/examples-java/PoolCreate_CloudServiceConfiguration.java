import com.azure.resourcemanager.batch.models.ApplicationPackageReference;
import com.azure.resourcemanager.batch.models.AutoUserScope;
import com.azure.resourcemanager.batch.models.AutoUserSpecification;
import com.azure.resourcemanager.batch.models.CertificateReference;
import com.azure.resourcemanager.batch.models.CertificateStoreLocation;
import com.azure.resourcemanager.batch.models.CertificateVisibility;
import com.azure.resourcemanager.batch.models.CloudServiceConfiguration;
import com.azure.resourcemanager.batch.models.ComputeNodeDeallocationOption;
import com.azure.resourcemanager.batch.models.ComputeNodeFillType;
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.ElevationLevel;
import com.azure.resourcemanager.batch.models.EnvironmentSetting;
import com.azure.resourcemanager.batch.models.FixedScaleSettings;
import com.azure.resourcemanager.batch.models.InterNodeCommunicationState;
import com.azure.resourcemanager.batch.models.IpAddressProvisioningType;
import com.azure.resourcemanager.batch.models.LinuxUserConfiguration;
import com.azure.resourcemanager.batch.models.MetadataItem;
import com.azure.resourcemanager.batch.models.NetworkConfiguration;
import com.azure.resourcemanager.batch.models.PublicIpAddressConfiguration;
import com.azure.resourcemanager.batch.models.ResourceFile;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import com.azure.resourcemanager.batch.models.StartTask;
import com.azure.resourcemanager.batch.models.TaskSchedulingPolicy;
import com.azure.resourcemanager.batch.models.UserAccount;
import com.azure.resourcemanager.batch.models.UserIdentity;
import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Pool Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolCreate_CloudServiceConfiguration.json
     */
    /**
     * Sample code: CreatePool - Full CloudServiceConfiguration.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void createPoolFullCloudServiceConfiguration(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .pools()
            .define("testpool")
            .withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withDisplayName("my-pool-name")
            .withVmSize("STANDARD_D4")
            .withDeploymentConfiguration(
                new DeploymentConfiguration()
                    .withCloudServiceConfiguration(
                        new CloudServiceConfiguration().withOsFamily("4").withOsVersion("WA-GUEST-OS-4.45_201708-01")))
            .withScaleSettings(
                new ScaleSettings()
                    .withFixedScale(
                        new FixedScaleSettings()
                            .withResizeTimeout(Duration.parse("PT8M"))
                            .withTargetDedicatedNodes(6)
                            .withTargetLowPriorityNodes(28)
                            .withNodeDeallocationOption(ComputeNodeDeallocationOption.TASK_COMPLETION)))
            .withInterNodeCommunication(InterNodeCommunicationState.ENABLED)
            .withNetworkConfiguration(
                new NetworkConfiguration()
                    .withSubnetId(
                        "/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123")
                    .withPublicIpAddressConfiguration(
                        new PublicIpAddressConfiguration()
                            .withProvision(IpAddressProvisioningType.USER_MANAGED)
                            .withIpAddressIds(
                                Arrays
                                    .asList(
                                        "/subscriptions/subid1/resourceGroups/rg13/providers/Microsoft.Network/publicIPAddresses/ip135",
                                        "/subscriptions/subid2/resourceGroups/rg24/providers/Microsoft.Network/publicIPAddresses/ip268"))))
            .withTaskSlotsPerNode(13)
            .withTaskSchedulingPolicy(new TaskSchedulingPolicy().withNodeFillType(ComputeNodeFillType.PACK))
            .withUserAccounts(
                Arrays
                    .asList(
                        new UserAccount()
                            .withName("username1")
                            .withPassword("fakeTokenPlaceholder")
                            .withElevationLevel(ElevationLevel.ADMIN)
                            .withLinuxUserConfiguration(
                                new LinuxUserConfiguration()
                                    .withUid(1234)
                                    .withGid(4567)
                                    .withSshPrivateKey("fakeTokenPlaceholder"))))
            .withMetadata(
                Arrays
                    .asList(
                        new MetadataItem().withName("metadata-1").withValue("value-1"),
                        new MetadataItem().withName("metadata-2").withValue("value-2")))
            .withStartTask(
                new StartTask()
                    .withCommandLine("cmd /c SET")
                    .withResourceFiles(
                        Arrays
                            .asList(
                                new ResourceFile()
                                    .withHttpUrl("https://testaccount.blob.core.windows.net/example-blob-file")
                                    .withFilePath("c:\\temp\\gohere")
                                    .withFileMode("777")))
                    .withEnvironmentSettings(
                        Arrays.asList(new EnvironmentSetting().withName("MYSET").withValue("1234")))
                    .withUserIdentity(
                        new UserIdentity()
                            .withAutoUser(
                                new AutoUserSpecification()
                                    .withScope(AutoUserScope.POOL)
                                    .withElevationLevel(ElevationLevel.ADMIN)))
                    .withMaxTaskRetryCount(6)
                    .withWaitForSuccess(true))
            .withCertificates(
                Arrays
                    .asList(
                        new CertificateReference()
                            .withId(
                                "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567")
                            .withStoreLocation(CertificateStoreLocation.LOCAL_MACHINE)
                            .withStoreName("MY")
                            .withVisibility(Arrays.asList(CertificateVisibility.REMOTE_USER))))
            .withApplicationPackages(
                Arrays
                    .asList(
                        new ApplicationPackageReference()
                            .withId(
                                "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234")
                            .withVersion("asdf")))
            .withApplicationLicenses(Arrays.asList("app-license0", "app-license1"))
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
