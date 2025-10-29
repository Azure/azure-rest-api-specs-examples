
import com.azure.resourcemanager.appcontainers.models.ActiveRevisionsMode;
import com.azure.resourcemanager.appcontainers.models.Configuration;
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerResources;
import com.azure.resourcemanager.appcontainers.models.EnvironmentVar;
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
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/
     * ContainerApps_Kind_FunctionApp_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update FunctionApp Kind.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateFunctionAppKind(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().define("testcontainerAppFunctionKind").withRegion("East Us")
            .withExistingResourceGroup("rg")
            .withManagedBy(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Web/sites/testcontainerAppFunctionKind")
            .withKind(Kind.FUNCTIONAPP)
            .withManagedEnvironmentId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3")
            .withConfiguration(new Configuration().withActiveRevisionsMode(ActiveRevisionsMode.SINGLE)
                .withIngress(new Ingress().withExternal(true).withTargetPort(80).withAllowInsecure(false)))
            .withTemplate(new Template().withContainers(Arrays.asList(new Container()
                .withImage("mcr.microsoft.com/azure-functions/dotnet:4").withName("function-app-container")
                .withEnv(Arrays.asList(new EnvironmentVar().withName("AzureWebJobsStorage").withValue(
                    "DefaultEndpointsProtocol=https;AccountName=mystorageaccount;AccountKey=mykey;EndpointSuffix=core.windows.net"),
                    new EnvironmentVar().withName("FUNCTIONS_WORKER_RUNTIME").withValue("dotnet"),
                    new EnvironmentVar().withName("WEBSITES_ENABLE_APP_SERVICE_STORAGE").withValue("false")))
                .withResources(new ContainerResources().withCpu(0.5D).withMemory("1.0Gi")))).withScale(
                    new Scale().withMinReplicas(0).withMaxReplicas(10).withCooldownPeriod(300).withPollingInterval(30)))
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
