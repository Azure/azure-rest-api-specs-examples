
import com.azure.resourcemanager.appcontainers.models.Action;
import com.azure.resourcemanager.appcontainers.models.Affinity;
import com.azure.resourcemanager.appcontainers.models.AppProtocol;
import com.azure.resourcemanager.appcontainers.models.BindingType;
import com.azure.resourcemanager.appcontainers.models.Configuration;
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbe;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGet;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGetHttpHeadersItem;
import com.azure.resourcemanager.appcontainers.models.ContainerResources;
import com.azure.resourcemanager.appcontainers.models.CorsPolicy;
import com.azure.resourcemanager.appcontainers.models.CustomDomain;
import com.azure.resourcemanager.appcontainers.models.CustomScaleRule;
import com.azure.resourcemanager.appcontainers.models.Dapr;
import com.azure.resourcemanager.appcontainers.models.ExtendedLocation;
import com.azure.resourcemanager.appcontainers.models.ExtendedLocationTypes;
import com.azure.resourcemanager.appcontainers.models.Ingress;
import com.azure.resourcemanager.appcontainers.models.IngressClientCertificateMode;
import com.azure.resourcemanager.appcontainers.models.IngressPortMapping;
import com.azure.resourcemanager.appcontainers.models.IngressStickySessions;
import com.azure.resourcemanager.appcontainers.models.InitContainer;
import com.azure.resourcemanager.appcontainers.models.IpSecurityRestrictionRule;
import com.azure.resourcemanager.appcontainers.models.Level;
import com.azure.resourcemanager.appcontainers.models.LogLevel;
import com.azure.resourcemanager.appcontainers.models.LoggerSetting;
import com.azure.resourcemanager.appcontainers.models.Runtime;
import com.azure.resourcemanager.appcontainers.models.RuntimeDotnet;
import com.azure.resourcemanager.appcontainers.models.RuntimeJava;
import com.azure.resourcemanager.appcontainers.models.RuntimeJavaAgent;
import com.azure.resourcemanager.appcontainers.models.RuntimeJavaAgentLogging;
import com.azure.resourcemanager.appcontainers.models.Scale;
import com.azure.resourcemanager.appcontainers.models.ScaleRule;
import com.azure.resourcemanager.appcontainers.models.Template;
import com.azure.resourcemanager.appcontainers.models.TrafficWeight;
import com.azure.resourcemanager.appcontainers.models.Type;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ContainerApps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ContainerApps_CreateOrUpdate_ConnectedEnvironment.json
     */
    /**
     * Sample code: Create or Update App On A Connected Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateAppOnAConnectedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().define("testcontainerApp0").withRegion("East US").withExistingResourceGroup("rg")
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation")
                .withType(ExtendedLocationTypes.CUSTOM_LOCATION))
            .withEnvironmentId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/connectedEnvironments/demokube")
            .withConfiguration(new Configuration().withIngress(new Ingress().withExternal(true).withTargetPort(3000)
                .withTraffic(Arrays.asList(new TrafficWeight().withRevisionName("testcontainerApp0-ab1234")
                    .withWeight(100).withLabel("production")))
                .withCustomDomains(Arrays.asList(new CustomDomain().withName("www.my-name.com")
                    .withBindingType(BindingType.SNI_ENABLED).withCertificateId(
                        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/connectedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
                    new CustomDomain().withName("www.my-other-name.com").withBindingType(BindingType.SNI_ENABLED)
                        .withCertificateId(
                            "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/connectedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com")))
                .withIpSecurityRestrictions(Arrays.asList(
                    new IpSecurityRestrictionRule().withName("Allow work IP A subnet")
                        .withDescription("Allowing all IP's within the subnet below to access containerapp")
                        .withIpAddressRange("192.168.1.1/32").withAction(Action.ALLOW),
                    new IpSecurityRestrictionRule().withName("Allow work IP B subnet")
                        .withDescription("Allowing all IP's within the subnet below to access containerapp")
                        .withIpAddressRange("192.168.1.1/8").withAction(Action.ALLOW)))
                .withStickySessions(new IngressStickySessions().withAffinity(Affinity.STICKY))
                .withClientCertificateMode(IngressClientCertificateMode.ACCEPT)
                .withCorsPolicy(new CorsPolicy()
                    .withAllowedOrigins(Arrays.asList("https://a.test.com", "https://b.test.com"))
                    .withAllowedMethods(Arrays.asList("GET", "POST"))
                    .withAllowedHeaders(Arrays.asList("HEADER1", "HEADER2"))
                    .withExposeHeaders(Arrays.asList("HEADER3", "HEADER4")).withMaxAge(1234).withAllowCredentials(true))
                .withAdditionalPortMappings(
                    Arrays.asList(new IngressPortMapping().withExternal(true).withTargetPort(1234),
                        new IngressPortMapping().withExternal(false).withTargetPort(2345).withExposedPort(3456))))
                .withDapr(new Dapr().withEnabled(true).withAppProtocol(AppProtocol.HTTP).withAppPort(3000)
                    .withHttpReadBufferSize(30).withHttpMaxRequestSize(10).withLogLevel(LogLevel.DEBUG)
                    .withEnableApiLogging(true))
                .withRuntime(new Runtime()
                    .withJava(new RuntimeJava().withEnableMetrics(true)
                        .withJavaAgent(new RuntimeJavaAgent().withEnabled(true)
                            .withLogging(new RuntimeJavaAgentLogging().withLoggerSettings(Arrays.asList(
                                new LoggerSetting().withLogger("org.springframework.boot").withLevel(Level.DEBUG))))))
                    .withDotnet(new RuntimeDotnet().withAutoConfigureDataProtection(true)))
                .withMaxInactiveRevisions(10))
            .withTemplate(new Template()
                .withInitContainers(Arrays.asList(new InitContainer().withImage("repo/testcontainerApp0:v4")
                    .withName("testinitcontainerApp0").withCommand(Arrays.asList("/bin/sh"))
                    .withArgs(Arrays.asList("-c", "while true; do echo hello; sleep 10;done"))
                    .withResources(new ContainerResources().withCpu(0.2D).withMemory("100Mi"))))
                .withContainers(
                    Arrays
                        .asList(new Container().withImage("repo/testcontainerApp0:v1").withName("testcontainerApp0")
                            .withProbes(Arrays.asList(new ContainerAppProbe()
                                .withHttpGet(new ContainerAppProbeHttpGet()
                                    .withHttpHeaders(Arrays.asList(new ContainerAppProbeHttpGetHttpHeadersItem()
                                        .withName("Custom-Header").withValue("Awesome")))
                                    .withPath("/health").withPort(8080))
                                .withInitialDelaySeconds(3).withPeriodSeconds(3).withType(Type.LIVENESS)))))
                .withScale(new Scale().withMinReplicas(1).withMaxReplicas(5)
                    .withRules(Arrays.asList(new ScaleRule().withName("httpscalingrule").withCustom(
                        new CustomScaleRule().withType("http").withMetadata(mapOf("concurrentRequests", "50")))))))
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
