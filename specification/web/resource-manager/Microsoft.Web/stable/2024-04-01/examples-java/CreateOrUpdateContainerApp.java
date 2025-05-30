
import com.azure.resourcemanager.appservice.fluent.models.ContainerAppInner;
import com.azure.resourcemanager.appservice.models.Configuration;
import com.azure.resourcemanager.appservice.models.Container;
import com.azure.resourcemanager.appservice.models.CustomScaleRule;
import com.azure.resourcemanager.appservice.models.Dapr;
import com.azure.resourcemanager.appservice.models.Ingress;
import com.azure.resourcemanager.appservice.models.Scale;
import com.azure.resourcemanager.appservice.models.ScaleRule;
import com.azure.resourcemanager.appservice.models.Template;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ContainerApps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/CreateOrUpdateContainerApp.json
     */
    /**
     * Sample code: Create or Update Container App.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateContainerApp(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getContainerApps().createOrUpdate("rg", "testcontainerApp0",
            new ContainerAppInner().withLocation("East US").withKind("containerApp").withKubeEnvironmentId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/kubeEnvironments/demokube")
                .withConfiguration(
                    new Configuration().withIngress(new Ingress().withExternal(true).withTargetPort(3000)))
                .withTemplate(new Template()
                    .withContainers(Arrays
                        .asList(new Container().withImage("repo/testcontainerApp0:v1").withName("testcontainerApp0")))
                    .withScale(new Scale().withMinReplicas(1).withMaxReplicas(5)
                        .withRules(Arrays.asList(new ScaleRule().withName("httpscalingrule").withCustom(
                            new CustomScaleRule().withType("http").withMetadata(mapOf("concurrentRequests", "50"))))))
                    .withDapr(new Dapr().withEnabled(true).withAppPort(3000))),
            com.azure.core.util.Context.NONE);
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
