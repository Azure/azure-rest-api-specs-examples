
import com.azure.resourcemanager.appcontainers.models.AppInsightsConfiguration;
import com.azure.resourcemanager.appcontainers.models.AppLogsConfiguration;
import com.azure.resourcemanager.appcontainers.models.CustomDomainConfiguration;
import com.azure.resourcemanager.appcontainers.models.DataDogConfiguration;
import com.azure.resourcemanager.appcontainers.models.DestinationsConfiguration;
import com.azure.resourcemanager.appcontainers.models.Header;
import com.azure.resourcemanager.appcontainers.models.LogAnalyticsConfiguration;
import com.azure.resourcemanager.appcontainers.models.LogsConfiguration;
import com.azure.resourcemanager.appcontainers.models.ManagedEnvironmentPropertiesPeerAuthentication;
import com.azure.resourcemanager.appcontainers.models.ManagedEnvironmentPropertiesPeerTrafficConfiguration;
import com.azure.resourcemanager.appcontainers.models.ManagedEnvironmentPropertiesPeerTrafficConfigurationEncryption;
import com.azure.resourcemanager.appcontainers.models.ManagedServiceIdentity;
import com.azure.resourcemanager.appcontainers.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.appcontainers.models.MetricsConfiguration;
import com.azure.resourcemanager.appcontainers.models.Mtls;
import com.azure.resourcemanager.appcontainers.models.OpenTelemetryConfiguration;
import com.azure.resourcemanager.appcontainers.models.OtlpConfiguration;
import com.azure.resourcemanager.appcontainers.models.TracesConfiguration;
import com.azure.resourcemanager.appcontainers.models.UserAssignedIdentity;
import com.azure.resourcemanager.appcontainers.models.VnetConfiguration;
import com.azure.resourcemanager.appcontainers.models.WorkloadProfile;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedEnvironments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/
     * ManagedEnvironments_CreateOrUpdate.json
     */
    /**
     * Sample code: Create environments.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createEnvironments(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironments().define("testcontainerenv").withRegion("East US")
            .withExistingResourceGroup("examplerg")
            .withIdentity(new ManagedServiceIdentity()
                .withType(ManagedServiceIdentityType.fromString("SystemAssigned, UserAssigned"))
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity",
                    new UserAssignedIdentity())))
            .withDaprAIConnectionString(
                "InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/")
            .withVnetConfiguration(new VnetConfiguration().withInfrastructureSubnetId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/RGName/providers/Microsoft.Network/virtualNetworks/VNetName/subnets/subnetName1"))
            .withAppLogsConfiguration(
                new AppLogsConfiguration().withLogAnalyticsConfiguration(new LogAnalyticsConfiguration()
                    .withCustomerId("string").withSharedKey("fakeTokenPlaceholder").withDynamicJsonColumns(true)))
            .withAppInsightsConfiguration(new AppInsightsConfiguration().withConnectionString(
                "InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://eastus-8.in.applicationinsights.azure.com/;LiveEndpoint=https://eastus.livediagnostics.monitor.azure.com/"))
            .withOpenTelemetryConfiguration(
                new OpenTelemetryConfiguration()
                    .withDestinationsConfiguration(
                        new DestinationsConfiguration()
                            .withDataDogConfiguration(
                                new DataDogConfiguration().withSite("string").withKey("fakeTokenPlaceholder"))
                            .withOtlpConfigurations(Arrays.asList(new OtlpConfiguration().withName("dashboard")
                                .withEndpoint("dashboard.k8s.region.azurecontainerapps.io:80").withInsecure(true)
                                .withHeaders(Arrays
                                    .asList(new Header().withKey("fakeTokenPlaceholder").withValue("xxxxxxxxxxx"))))))
                    .withTracesConfiguration(
                        new TracesConfiguration().withIncludeDapr(true).withDestinations(Arrays.asList("appInsights")))
                    .withLogsConfiguration(new LogsConfiguration().withDestinations(Arrays.asList("appInsights")))
                    .withMetricsConfiguration(
                        new MetricsConfiguration().withIncludeKeda(true).withDestinations(Arrays.asList("dataDog"))))
            .withZoneRedundant(true)
            .withCustomDomainConfiguration(new CustomDomainConfiguration().withDnsSuffix("www.my-name.com")
                .withCertificateValue("Y2VydA==".getBytes()).withCertificatePassword("fakeTokenPlaceholder"))
            .withWorkloadProfiles(Arrays.asList(
                new WorkloadProfile().withName("My-GP-01").withEnableFips(true)
                    .withWorkloadProfileType("GeneralPurpose").withMinimumCount(3).withMaximumCount(12),
                new WorkloadProfile().withName("My-MO-01").withWorkloadProfileType("MemoryOptimized")
                    .withMinimumCount(3).withMaximumCount(6),
                new WorkloadProfile().withName("My-CO-01").withWorkloadProfileType("ComputeOptimized")
                    .withMinimumCount(3).withMaximumCount(6),
                new WorkloadProfile().withName("My-consumption-01").withWorkloadProfileType("Consumption")))
            .withPeerAuthentication(
                new ManagedEnvironmentPropertiesPeerAuthentication().withMtls(new Mtls().withEnabled(true)))
            .withPeerTrafficConfiguration(new ManagedEnvironmentPropertiesPeerTrafficConfiguration()
                .withEncryption(new ManagedEnvironmentPropertiesPeerTrafficConfigurationEncryption().withEnabled(true)))
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
