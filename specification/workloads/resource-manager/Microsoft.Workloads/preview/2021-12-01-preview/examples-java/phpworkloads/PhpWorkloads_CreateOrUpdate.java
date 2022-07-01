import com.azure.resourcemanager.workloads.models.AzureFrontDoorEnabled;
import com.azure.resourcemanager.workloads.models.BackupProfile;
import com.azure.resourcemanager.workloads.models.CacheProfile;
import com.azure.resourcemanager.workloads.models.DatabaseProfile;
import com.azure.resourcemanager.workloads.models.DatabaseTier;
import com.azure.resourcemanager.workloads.models.DatabaseType;
import com.azure.resourcemanager.workloads.models.DiskInfo;
import com.azure.resourcemanager.workloads.models.DiskStorageType;
import com.azure.resourcemanager.workloads.models.EnableBackup;
import com.azure.resourcemanager.workloads.models.EnableSslEnforcement;
import com.azure.resourcemanager.workloads.models.FileShareStorageType;
import com.azure.resourcemanager.workloads.models.FileShareType;
import com.azure.resourcemanager.workloads.models.FileshareProfile;
import com.azure.resourcemanager.workloads.models.HAEnabled;
import com.azure.resourcemanager.workloads.models.LoadBalancerType;
import com.azure.resourcemanager.workloads.models.ManagedRGConfiguration;
import com.azure.resourcemanager.workloads.models.NetworkProfile;
import com.azure.resourcemanager.workloads.models.NodeProfile;
import com.azure.resourcemanager.workloads.models.OSImageOffer;
import com.azure.resourcemanager.workloads.models.OSImagePublisher;
import com.azure.resourcemanager.workloads.models.OSImageSku;
import com.azure.resourcemanager.workloads.models.OSImageVersion;
import com.azure.resourcemanager.workloads.models.OsImageProfile;
import com.azure.resourcemanager.workloads.models.PhpProfile;
import com.azure.resourcemanager.workloads.models.PhpVersion;
import com.azure.resourcemanager.workloads.models.RedisCacheFamily;
import com.azure.resourcemanager.workloads.models.SearchProfile;
import com.azure.resourcemanager.workloads.models.SearchType;
import com.azure.resourcemanager.workloads.models.SiteProfile;
import com.azure.resourcemanager.workloads.models.Sku;
import com.azure.resourcemanager.workloads.models.UserProfile;
import com.azure.resourcemanager.workloads.models.VmssNodesProfile;
import com.azure.resourcemanager.workloads.models.WorkloadKind;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for PhpWorkloads CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/phpworkloads/PhpWorkloads_CreateOrUpdate.json
     */
    /**
     * Sample code: Workloads.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void workloads(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .phpWorkloads()
            .define("wp39")
            .withRegion("eastus2")
            .withExistingResourceGroup("test-rg")
            .withKind(WorkloadKind.WORD_PRESS)
            .withTags(mapOf())
            .withSku(new Sku().withName("Large"))
            .withAppLocation("eastus")
            .withManagedResourceGroupConfiguration(new ManagedRGConfiguration().withName("php-mrg-wp39"))
            .withAdminUserProfile(new UserProfile().withUsername("wpadmin").withSshPublicKey("===SSH=PUBLIC=KEY==="))
            .withWebNodesProfile(
                new VmssNodesProfile()
                    .withName("web-server")
                    .withNodeSku("Standard_DS2_v2")
                    .withOsImage(
                        new OsImageProfile()
                            .withPublisher(OSImagePublisher.CANONICAL)
                            .withOffer(OSImageOffer.UBUNTU_SERVER)
                            .withSku(OSImageSku.fromString("18.0-LTS"))
                            .withVersion(OSImageVersion.LATEST))
                    .withOsDisk(new DiskInfo().withStorageType(DiskStorageType.PREMIUM_LRS))
                    .withAutoScaleMinCount(1)
                    .withAutoScaleMaxCount(1))
            .withControllerProfile(
                new NodeProfile()
                    .withName("contoller-vm")
                    .withNodeSku("Standard_DS2_v2")
                    .withOsImage(
                        new OsImageProfile()
                            .withPublisher(OSImagePublisher.CANONICAL)
                            .withOffer(OSImageOffer.UBUNTU_SERVER)
                            .withSku(OSImageSku.fromString("18.0-LTS"))
                            .withVersion(OSImageVersion.LATEST))
                    .withOsDisk(new DiskInfo().withStorageType(DiskStorageType.PREMIUM_LRS))
                    .withDataDisks(
                        Arrays.asList(new DiskInfo().withStorageType(DiskStorageType.PREMIUM_LRS).withSizeInGB(100L))))
            .withNetworkProfile(
                new NetworkProfile()
                    .withLoadBalancerType(LoadBalancerType.LOAD_BALANCER)
                    .withLoadBalancerSku("Standard")
                    .withAzureFrontDoorEnabled(AzureFrontDoorEnabled.ENABLED))
            .withDatabaseProfile(
                new DatabaseProfile()
                    .withType(DatabaseType.MY_SQL)
                    .withServerName("wp-db-server")
                    .withVersion("5.7")
                    .withSku("Standard_D32s_v4")
                    .withTier(DatabaseTier.GENERAL_PURPOSE)
                    .withHaEnabled(HAEnabled.DISABLED)
                    .withStorageSku("Premium_LRS")
                    .withStorageInGB(128L)
                    .withStorageIops(200L)
                    .withBackupRetentionDays(7)
                    .withSslEnforcementEnabled(EnableSslEnforcement.ENABLED))
            .withSiteProfile(new SiteProfile().withDomainName("www.example.com"))
            .withFileshareProfile(
                new FileshareProfile()
                    .withShareType(FileShareType.AZURE_FILES)
                    .withStorageType(FileShareStorageType.PREMIUM_LRS)
                    .withShareSizeInGB(100L))
            .withPhpProfile(new PhpProfile().withVersion(PhpVersion.SEVEN_THREE))
            .withSearchProfile(
                new SearchProfile()
                    .withNodeSku("Standard_DS2_v2")
                    .withOsImage(
                        new OsImageProfile()
                            .withPublisher(OSImagePublisher.CANONICAL)
                            .withOffer(OSImageOffer.UBUNTU_SERVER)
                            .withSku(OSImageSku.fromString("18.0-LTS"))
                            .withVersion(OSImageVersion.LATEST))
                    .withOsDisk(new DiskInfo().withStorageType(DiskStorageType.PREMIUM_LRS))
                    .withSearchType(SearchType.ELASTIC))
            .withCacheProfile(
                new CacheProfile()
                    .withName("wp-cache")
                    .withSkuName("Basic")
                    .withFamily(RedisCacheFamily.C)
                    .withCapacity(0L))
            .withBackupProfile(new BackupProfile().withBackupEnabled(EnableBackup.DISABLED))
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
