import com.azure.resourcemanager.appcontainers.models.AppProtocol;
import com.azure.resourcemanager.appcontainers.models.BindingType;
import com.azure.resourcemanager.appcontainers.models.Configuration;
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbe;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGet;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGetHttpHeadersItem;
import com.azure.resourcemanager.appcontainers.models.CustomDomain;
import com.azure.resourcemanager.appcontainers.models.CustomScaleRule;
import com.azure.resourcemanager.appcontainers.models.Dapr;
import com.azure.resourcemanager.appcontainers.models.Ingress;
import com.azure.resourcemanager.appcontainers.models.Scale;
import com.azure.resourcemanager.appcontainers.models.ScaleRule;
import com.azure.resourcemanager.appcontainers.models.Template;
import com.azure.resourcemanager.appcontainers.models.TrafficWeight;
import com.azure.resourcemanager.appcontainers.models.Type;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ContainerApps CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ContainerApps_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Container App.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateContainerApp(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerApps()
            .define("testcontainerApp0")
            .withRegion("East US")
            .withExistingResourceGroup("rg")
            .withManagedEnvironmentId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube")
            .withConfiguration(
                new Configuration()
                    .withIngress(
                        new Ingress()
                            .withExternal(true)
                            .withTargetPort(3000)
                            .withTraffic(
                                Arrays
                                    .asList(
                                        new TrafficWeight()
                                            .withRevisionName("testcontainerApp0-ab1234")
                                            .withWeight(100)
                                            .withLabel("production")))
                            .withCustomDomains(
                                Arrays
                                    .asList(
                                        new CustomDomain()
                                            .withName("www.my-name.com")
                                            .withBindingType(BindingType.SNI_ENABLED)
                                            .withCertificateId(
                                                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
                                        new CustomDomain()
                                            .withName("www.my-other-name.com")
                                            .withBindingType(BindingType.SNI_ENABLED)
                                            .withCertificateId(
                                                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com"))))
                    .withDapr(new Dapr().withEnabled(true).withAppProtocol(AppProtocol.HTTP).withAppPort(3000)))
            .withTemplate(
                new Template()
                    .withContainers(
                        Arrays
                            .asList(
                                new Container()
                                    .withImage("repo/testcontainerApp0:v1")
                                    .withName("testcontainerApp0")
                                    .withProbes(
                                        Arrays
                                            .asList(
                                                new ContainerAppProbe()
                                                    .withHttpGet(
                                                        new ContainerAppProbeHttpGet()
                                                            .withHttpHeaders(
                                                                Arrays
                                                                    .asList(
                                                                        new ContainerAppProbeHttpGetHttpHeadersItem()
                                                                            .withName("Custom-Header")
                                                                            .withValue("Awesome")))
                                                            .withPath("/health")
                                                            .withPort(8080))
                                                    .withInitialDelaySeconds(3)
                                                    .withPeriodSeconds(3)
                                                    .withType(Type.LIVENESS)))))
                    .withScale(
                        new Scale()
                            .withMinReplicas(1)
                            .withMaxReplicas(5)
                            .withRules(
                                Arrays
                                    .asList(
                                        new ScaleRule()
                                            .withName("httpscalingrule")
                                            .withCustom(
                                                new CustomScaleRule()
                                                    .withType("http")
                                                    .withMetadata(mapOf("concurrentRequests", "50")))))))
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
