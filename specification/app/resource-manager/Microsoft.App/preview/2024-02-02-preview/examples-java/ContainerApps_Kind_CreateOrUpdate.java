
import com.azure.resourcemanager.appcontainers.models.ActiveRevisionsMode;
import com.azure.resourcemanager.appcontainers.models.Configuration;
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerResources;
import com.azure.resourcemanager.appcontainers.models.Ingress;
import com.azure.resourcemanager.appcontainers.models.Kind;
import com.azure.resourcemanager.appcontainers.models.Scale;
import com.azure.resourcemanager.appcontainers.models.Template;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ContainerApps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ContainerApps_Kind_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update App Kind.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateAppKind(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().define("testcontainerAppKind").withRegion("East Us").withExistingResourceGroup("rg")
            .withManagedBy(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Web/sites/testcontainerAppKind")
            .withKind(Kind.WORKFLOWAPP)
            .withManagedEnvironmentId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3")
            .withConfiguration(new Configuration().withActiveRevisionsMode(ActiveRevisionsMode.SINGLE)
                .withIngress(new Ingress().withExternal(true).withTargetPort(80).withAllowInsecure(true)))
            .withTemplate(new Template()
                .withContainers(Arrays
                    .asList(new Container().withImage("default/logicapps-base:latest").withName("logicapps-container")
                        .withResources(new ContainerResources().withCpu(1.0D).withMemory("2.0Gi"))))
                .withScale(new Scale().withMinReplicas(1).withMaxReplicas(30)))
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
