
import com.azure.resourcemanager.servicefabric.models.AddOnFeatures;
import com.azure.resourcemanager.servicefabric.models.ApplicationDeltaHealthPolicy;
import com.azure.resourcemanager.servicefabric.models.ApplicationHealthPolicy;
import com.azure.resourcemanager.servicefabric.models.ApplicationTypeVersionsCleanupPolicy;
import com.azure.resourcemanager.servicefabric.models.AzureActiveDirectory;
import com.azure.resourcemanager.servicefabric.models.ClientCertificateCommonName;
import com.azure.resourcemanager.servicefabric.models.ClientCertificateThumbprint;
import com.azure.resourcemanager.servicefabric.models.ClusterHealthPolicy;
import com.azure.resourcemanager.servicefabric.models.ClusterUpgradeCadence;
import com.azure.resourcemanager.servicefabric.models.ClusterUpgradeDeltaHealthPolicy;
import com.azure.resourcemanager.servicefabric.models.ClusterUpgradePolicy;
import com.azure.resourcemanager.servicefabric.models.DiagnosticsStorageAccountConfig;
import com.azure.resourcemanager.servicefabric.models.DurabilityLevel;
import com.azure.resourcemanager.servicefabric.models.EndpointRangeDescription;
import com.azure.resourcemanager.servicefabric.models.NodeTypeDescription;
import com.azure.resourcemanager.servicefabric.models.Notification;
import com.azure.resourcemanager.servicefabric.models.NotificationCategory;
import com.azure.resourcemanager.servicefabric.models.NotificationChannel;
import com.azure.resourcemanager.servicefabric.models.NotificationLevel;
import com.azure.resourcemanager.servicefabric.models.NotificationTarget;
import com.azure.resourcemanager.servicefabric.models.ReliabilityLevel;
import com.azure.resourcemanager.servicefabric.models.ServerCertificateCommonName;
import com.azure.resourcemanager.servicefabric.models.ServerCertificateCommonNames;
import com.azure.resourcemanager.servicefabric.models.ServiceTypeDeltaHealthPolicy;
import com.azure.resourcemanager.servicefabric.models.ServiceTypeHealthPolicy;
import com.azure.resourcemanager.servicefabric.models.SettingsParameterDescription;
import com.azure.resourcemanager.servicefabric.models.SettingsSectionDescription;
import com.azure.resourcemanager.servicefabric.models.SfZonalUpgradeMode;
import com.azure.resourcemanager.servicefabric.models.StoreName;
import com.azure.resourcemanager.servicefabric.models.UpgradeMode;
import com.azure.resourcemanager.servicefabric.models.VmssZonalUpgradeMode;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/
     * ClusterPutOperation_example_max.json
     */
    /**
     * Sample code: Put a cluster with maximum parameters.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void
        putAClusterWithMaximumParameters(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.clusters().define("myCluster").withRegion("eastus").withExistingResourceGroup("resRg").withTags(mapOf())
            .withAddOnFeatures(Arrays.asList(AddOnFeatures.REPAIR_MANAGER, AddOnFeatures.DNS_SERVICE,
                AddOnFeatures.BACKUP_RESTORE_SERVICE, AddOnFeatures.RESOURCE_MONITOR_SERVICE))
            .withAzureActiveDirectory(new AzureActiveDirectory().withTenantId("6abcc6a0-8666-43f1-87b8-172cf86a9f9c")
                .withClusterApplication("5886372e-7bf4-4878-a497-8098aba608ae")
                .withClientApplication("d151ad89-4bce-4ae8-b3d1-1dc79679fa75"))
            .withCertificateCommonNames(new ServerCertificateCommonNames()
                .withCommonNames(Arrays.asList(new ServerCertificateCommonName().withCertificateCommonName("abc.com")
                    .withCertificateIssuerThumbprint("12599211F8F14C90AFA9532AD79A6F2CA1C00622")))
                .withX509StoreName(StoreName.MY))
            .withClientCertificateCommonNames(
                Arrays.asList(new ClientCertificateCommonName().withIsAdmin(true).withCertificateCommonName("abc.com")
                    .withCertificateIssuerThumbprint("5F3660C715EBBDA31DB1FFDCF508302348DE8E7A")))
            .withClientCertificateThumbprints(Arrays.asList(new ClientCertificateThumbprint().withIsAdmin(true)
                .withCertificateThumbprint("5F3660C715EBBDA31DB1FFDCF508302348DE8E7A")))
            .withClusterCodeVersion("7.0.470.9590")
            .withDiagnosticsStorageAccountConfig(new DiagnosticsStorageAccountConfig().withStorageAccountName("diag")
                .withProtectedAccountKeyName("fakeTokenPlaceholder")
                .withBlobEndpoint("https://diag.blob.core.windows.net/")
                .withQueueEndpoint("https://diag.queue.core.windows.net/")
                .withTableEndpoint("https://diag.table.core.windows.net/"))
            .withEventStoreServiceEnabled(true)
            .withFabricSettings(Arrays.asList(new SettingsSectionDescription().withName("UpgradeService")
                .withParameters(Arrays
                    .asList(new SettingsParameterDescription().withName("AppPollIntervalInSeconds").withValue("60")))))
            .withManagementEndpoint("https://myCluster.eastus.cloudapp.azure.com:19080")
            .withNodeTypes(Arrays.asList(new NodeTypeDescription().withName("nt1vm")
                .withClientConnectionEndpointPort(19000).withHttpGatewayEndpointPort(19007)
                .withDurabilityLevel(DurabilityLevel.SILVER)
                .withApplicationPorts(new EndpointRangeDescription().withStartPort(20000).withEndPort(30000))
                .withEphemeralPorts(new EndpointRangeDescription().withStartPort(49000).withEndPort(64000))
                .withIsPrimary(true).withVmInstanceCount(5).withIsStateless(false).withMultipleAvailabilityZones(true)))
            .withReliabilityLevel(ReliabilityLevel.PLATINUM)
            .withReverseProxyCertificateCommonNames(new ServerCertificateCommonNames()
                .withCommonNames(Arrays.asList(new ServerCertificateCommonName().withCertificateCommonName("abc.com")
                    .withCertificateIssuerThumbprint("12599211F8F14C90AFA9532AD79A6F2CA1C00622")))
                .withX509StoreName(StoreName.MY))
            .withUpgradeDescription(new ClusterUpgradePolicy().withForceRestart(false)
                .withUpgradeReplicaSetCheckTimeout("00:10:00").withHealthCheckWaitDuration("00:00:30")
                .withHealthCheckStableDuration("00:00:30").withHealthCheckRetryTimeout("00:05:00")
                .withUpgradeTimeout("01:00:00").withUpgradeDomainTimeout("00:15:00")
                .withHealthPolicy(
                    new ClusterHealthPolicy().withMaxPercentUnhealthyNodes(0).withMaxPercentUnhealthyApplications(0)
                        .withApplicationHealthPolicies(mapOf("fabric:/myApp1",
                            new ApplicationHealthPolicy()
                                .withDefaultServiceTypeHealthPolicy(
                                    new ServiceTypeHealthPolicy().withMaxPercentUnhealthyServices(0))
                                .withServiceTypeHealthPolicies(mapOf("myServiceType1",
                                    new ServiceTypeHealthPolicy().withMaxPercentUnhealthyServices(100))))))
                .withDeltaHealthPolicy(new ClusterUpgradeDeltaHealthPolicy().withMaxPercentDeltaUnhealthyNodes(0)
                    .withMaxPercentUpgradeDomainDeltaUnhealthyNodes(0).withMaxPercentDeltaUnhealthyApplications(0)
                    .withApplicationDeltaHealthPolicies(mapOf("fabric:/myApp1",
                        new ApplicationDeltaHealthPolicy()
                            .withDefaultServiceTypeDeltaHealthPolicy(
                                new ServiceTypeDeltaHealthPolicy().withMaxPercentDeltaUnhealthyServices(0))
                            .withServiceTypeDeltaHealthPolicies(mapOf("myServiceType1",
                                new ServiceTypeDeltaHealthPolicy().withMaxPercentDeltaUnhealthyServices(0)))))))
            .withUpgradeMode(UpgradeMode.MANUAL)
            .withApplicationTypeVersionsCleanupPolicy(
                new ApplicationTypeVersionsCleanupPolicy().withMaxUnusedVersionsToKeep(2L))
            .withVmImage("Windows").withSfZonalUpgradeMode(SfZonalUpgradeMode.HIERARCHICAL)
            .withVmssZonalUpgradeMode(VmssZonalUpgradeMode.PARALLEL).withInfrastructureServiceManager(true)
            .withUpgradeWave(ClusterUpgradeCadence.WAVE1)
            .withUpgradePauseStartTimestampUtc(OffsetDateTime.parse("2021-06-21T22:00:00Z"))
            .withUpgradePauseEndTimestampUtc(OffsetDateTime.parse("2021-06-25T22:00:00Z"))
            .withNotifications(Arrays.asList(
                new Notification().withIsEnabled(true).withNotificationCategory(NotificationCategory.WAVE_PROGRESS)
                    .withNotificationLevel(NotificationLevel.CRITICAL)
                    .withNotificationTargets(Arrays.asList(
                        new NotificationTarget().withNotificationChannel(NotificationChannel.EMAIL_USER)
                            .withReceivers(Arrays.asList("****@microsoft.com", "****@microsoft.com")),
                        new NotificationTarget().withNotificationChannel(NotificationChannel.EMAIL_SUBSCRIPTION)
                            .withReceivers(Arrays.asList("Owner", "AccountAdmin")))),
                new Notification().withIsEnabled(true).withNotificationCategory(NotificationCategory.WAVE_PROGRESS)
                    .withNotificationLevel(NotificationLevel.ALL)
                    .withNotificationTargets(Arrays.asList(
                        new NotificationTarget().withNotificationChannel(NotificationChannel.EMAIL_USER)
                            .withReceivers(Arrays.asList("****@microsoft.com", "****@microsoft.com")),
                        new NotificationTarget().withNotificationChannel(NotificationChannel.EMAIL_SUBSCRIPTION)
                            .withReceivers(Arrays.asList("Owner", "AccountAdmin"))))))
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
