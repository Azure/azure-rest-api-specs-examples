import com.azure.resourcemanager.appcontainers.models.Configuration;
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbe;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeTcpSocket;
import com.azure.resourcemanager.appcontainers.models.Ingress;
import com.azure.resourcemanager.appcontainers.models.IngressTransportMethod;
import com.azure.resourcemanager.appcontainers.models.Scale;
import com.azure.resourcemanager.appcontainers.models.ScaleRule;
import com.azure.resourcemanager.appcontainers.models.TcpScaleRule;
import com.azure.resourcemanager.appcontainers.models.Template;
import com.azure.resourcemanager.appcontainers.models.TrafficWeight;
import com.azure.resourcemanager.appcontainers.models.Type;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ContainerApps CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ContainerApps_ManagedBy_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update ManagedBy App.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateManagedByApp(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerApps()
            .define("testcontainerAppManagedBy")
            .withRegion("East US")
            .withExistingResourceGroup("rg")
            .withManagedBy(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.AppPlatform/Spring/springapp")
            .withEnvironmentId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube")
            .withConfiguration(
                new Configuration()
                    .withIngress(
                        new Ingress()
                            .withExternal(true)
                            .withTargetPort(3000)
                            .withExposedPort(4000)
                            .withTransport(IngressTransportMethod.TCP)
                            .withTraffic(
                                Arrays
                                    .asList(
                                        new TrafficWeight()
                                            .withRevisionName("testcontainerAppManagedBy-ab1234")
                                            .withWeight(100)))))
            .withTemplate(
                new Template()
                    .withContainers(
                        Arrays
                            .asList(
                                new Container()
                                    .withImage("repo/testcontainerAppManagedBy:v1")
                                    .withName("testcontainerAppManagedBy")
                                    .withProbes(
                                        Arrays
                                            .asList(
                                                new ContainerAppProbe()
                                                    .withInitialDelaySeconds(3)
                                                    .withPeriodSeconds(3)
                                                    .withTcpSocket(new ContainerAppProbeTcpSocket().withPort(8080))
                                                    .withType(Type.LIVENESS)))))
                    .withScale(
                        new Scale()
                            .withMinReplicas(1)
                            .withMaxReplicas(5)
                            .withRules(
                                Arrays
                                    .asList(
                                        new ScaleRule()
                                            .withName("tcpscalingrule")
                                            .withTcp(
                                                new TcpScaleRule()
                                                    .withMetadata(mapOf("concurrentConnections", "50")))))))
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
