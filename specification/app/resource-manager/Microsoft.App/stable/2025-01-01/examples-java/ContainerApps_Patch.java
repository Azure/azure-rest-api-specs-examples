
import com.azure.resourcemanager.appcontainers.models.Action;
import com.azure.resourcemanager.appcontainers.models.Affinity;
import com.azure.resourcemanager.appcontainers.models.AppProtocol;
import com.azure.resourcemanager.appcontainers.models.BindingType;
import com.azure.resourcemanager.appcontainers.models.Configuration;
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerApp;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbe;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGet;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGetHttpHeadersItem;
import com.azure.resourcemanager.appcontainers.models.ContainerResources;
import com.azure.resourcemanager.appcontainers.models.CustomDomain;
import com.azure.resourcemanager.appcontainers.models.CustomScaleRule;
import com.azure.resourcemanager.appcontainers.models.Dapr;
import com.azure.resourcemanager.appcontainers.models.Ingress;
import com.azure.resourcemanager.appcontainers.models.IngressStickySessions;
import com.azure.resourcemanager.appcontainers.models.InitContainer;
import com.azure.resourcemanager.appcontainers.models.IpSecurityRestrictionRule;
import com.azure.resourcemanager.appcontainers.models.LogLevel;
import com.azure.resourcemanager.appcontainers.models.Runtime;
import com.azure.resourcemanager.appcontainers.models.RuntimeJava;
import com.azure.resourcemanager.appcontainers.models.Scale;
import com.azure.resourcemanager.appcontainers.models.ScaleRule;
import com.azure.resourcemanager.appcontainers.models.Service;
import com.azure.resourcemanager.appcontainers.models.ServiceBind;
import com.azure.resourcemanager.appcontainers.models.Template;
import com.azure.resourcemanager.appcontainers.models.TrafficWeight;
import com.azure.resourcemanager.appcontainers.models.Type;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ContainerApps Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ContainerApps_Patch.json
     */
    /**
     * Sample code: Patch Container App.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void patchContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        ContainerApp resource = manager.containerApps()
            .getByResourceGroupWithResponse("rg", "testcontainerapp0", com.azure.core.util.Context.NONE).getValue();
        resource
            .update().withTags(
                mapOf("tag1", "value1", "tag2", "value2"))
            .withConfiguration(new Configuration().withIngress(new Ingress().withExternal(true).withTargetPort(3000)
                .withTraffic(Arrays.asList(new TrafficWeight().withRevisionName("testcontainerapp0-ab1234")
                    .withWeight(100).withLabel("production")))
                .withCustomDomains(Arrays.asList(new CustomDomain().withName("www.my-name.com")
                    .withBindingType(BindingType.SNI_ENABLED).withCertificateId(
                        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
                    new CustomDomain().withName("www.my-other-name.com").withBindingType(BindingType.SNI_ENABLED)
                        .withCertificateId(
                            "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com")))
                .withIpSecurityRestrictions(Arrays.asList(
                    new IpSecurityRestrictionRule().withName("Allow work IP A subnet")
                        .withDescription("Allowing all IP's within the subnet below to access containerapp")
                        .withIpAddressRange("192.168.1.1/32").withAction(Action.ALLOW),
                    new IpSecurityRestrictionRule().withName("Allow work IP B subnet")
                        .withDescription("Allowing all IP's within the subnet below to access containerapp")
                        .withIpAddressRange("192.168.1.1/8").withAction(Action.ALLOW)))
                .withStickySessions(new IngressStickySessions().withAffinity(Affinity.STICKY)))
                .withDapr(new Dapr().withEnabled(true).withAppProtocol(AppProtocol.HTTP).withAppPort(3000)
                    .withHttpReadBufferSize(30).withHttpMaxRequestSize(10).withLogLevel(LogLevel.DEBUG)
                    .withEnableApiLogging(true))
                .withRuntime(new Runtime().withJava(new RuntimeJava().withEnableMetrics(true)))
                .withMaxInactiveRevisions(10).withService(new Service().withType("redis")))
            .withTemplate(
                new Template()
                    .withInitContainers(Arrays.asList(
                        new InitContainer().withImage("repo/testcontainerapp0:v4").withName("testinitcontainerApp0")
                            .withResources(new ContainerResources().withCpu(0.5D).withMemory("1Gi"))))
                    .withContainers(Arrays.asList(new Container().withImage("repo/testcontainerapp0:v1")
                        .withName("testcontainerapp0").withProbes(Arrays.asList(new ContainerAppProbe()
                            .withHttpGet(new ContainerAppProbeHttpGet()
                                .withHttpHeaders(Arrays.asList(new ContainerAppProbeHttpGetHttpHeadersItem()
                                    .withName("Custom-Header").withValue("Awesome")))
                                .withPath("/health").withPort(8080))
                            .withInitialDelaySeconds(3).withPeriodSeconds(3).withType(Type.LIVENESS)))))
                    .withScale(new Scale()
                        .withMinReplicas(1).withMaxReplicas(5).withCooldownPeriod(350).withPollingInterval(35)
                        .withRules(
                            Arrays.asList(new ScaleRule().withName("httpscalingrule")
                                .withCustom(new CustomScaleRule().withType("http")
                                    .withMetadata(mapOf("concurrentRequests", "50"))))))
                    .withServiceBinds(Arrays.asList(new ServiceBind().withServiceId(
                        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/service")
                        .withName("service"))))
            .apply();
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
