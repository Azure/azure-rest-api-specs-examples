import com.azure.resourcemanager.appcontainers.models.AppLogsConfiguration;
import com.azure.resourcemanager.appcontainers.models.CustomDomainConfiguration;
import com.azure.resourcemanager.appcontainers.models.EnvironmentSkuProperties;
import com.azure.resourcemanager.appcontainers.models.LogAnalyticsConfiguration;
import com.azure.resourcemanager.appcontainers.models.ManagedEnvironmentOutBoundType;
import com.azure.resourcemanager.appcontainers.models.ManagedEnvironmentOutboundSettings;
import com.azure.resourcemanager.appcontainers.models.SkuName;
import com.azure.resourcemanager.appcontainers.models.VnetConfiguration;
import com.azure.resourcemanager.appcontainers.models.WorkloadProfile;
import java.util.Arrays;

/** Samples for ManagedEnvironments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ManagedEnvironments_CreateOrUpdate.json
     */
    /**
     * Sample code: Create environments.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createEnvironments(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .managedEnvironments()
            .define("testcontainerenv")
            .withRegion("East US")
            .withExistingResourceGroup("examplerg")
            .withSku(new EnvironmentSkuProperties().withName(SkuName.PREMIUM))
            .withDaprAIConnectionString(
                "InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/")
            .withVnetConfiguration(
                new VnetConfiguration()
                    .withOutboundSettings(
                        new ManagedEnvironmentOutboundSettings()
                            .withOutBoundType(ManagedEnvironmentOutBoundType.USER_DEFINED_ROUTING)
                            .withVirtualNetworkApplianceIp("192.168.1.20")))
            .withAppLogsConfiguration(
                new AppLogsConfiguration()
                    .withLogAnalyticsConfiguration(
                        new LogAnalyticsConfiguration().withCustomerId("string").withSharedKey("string")))
            .withZoneRedundant(true)
            .withCustomDomainConfiguration(
                new CustomDomainConfiguration()
                    .withDnsSuffix("www.my-name.com")
                    .withCertificateValue("PFX-or-PEM-blob".getBytes())
                    .withCertificatePassword("private key password".getBytes()))
            .withWorkloadProfiles(
                Arrays
                    .asList(
                        new WorkloadProfile()
                            .withWorkloadProfileType("GeneralPurpose")
                            .withMinimumCount(3)
                            .withMaximumCount(12),
                        new WorkloadProfile()
                            .withWorkloadProfileType("MemoryOptimized")
                            .withMinimumCount(3)
                            .withMaximumCount(6),
                        new WorkloadProfile()
                            .withWorkloadProfileType("ComputeOptimized")
                            .withMinimumCount(3)
                            .withMaximumCount(6)))
            .create();
    }
}
