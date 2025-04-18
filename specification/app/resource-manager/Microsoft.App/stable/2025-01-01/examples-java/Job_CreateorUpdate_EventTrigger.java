
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerResources;
import com.azure.resourcemanager.appcontainers.models.InitContainer;
import com.azure.resourcemanager.appcontainers.models.JobConfiguration;
import com.azure.resourcemanager.appcontainers.models.JobConfigurationEventTriggerConfig;
import com.azure.resourcemanager.appcontainers.models.JobScale;
import com.azure.resourcemanager.appcontainers.models.JobScaleRule;
import com.azure.resourcemanager.appcontainers.models.JobTemplate;
import com.azure.resourcemanager.appcontainers.models.ManagedServiceIdentity;
import com.azure.resourcemanager.appcontainers.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.appcontainers.models.TriggerType;
import com.azure.resourcemanager.appcontainers.models.UserAssignedIdentity;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Jobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Job_CreateorUpdate_EventTrigger.json
     */
    /**
     * Sample code: Create or Update Container Apps Job With Event Driven Trigger.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateContainerAppsJobWithEventDrivenTrigger(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) throws IOException {
        manager.jobs().define("testcontainerappsjob0").withRegion("East US").withExistingResourceGroup("rg")
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity",
                    new UserAssignedIdentity())))
            .withEnvironmentId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube")
            .withConfiguration(new JobConfiguration().withTriggerType(TriggerType.EVENT).withReplicaTimeout(10)
                .withReplicaRetryLimit(10)
                .withEventTriggerConfig(new JobConfigurationEventTriggerConfig().withReplicaCompletionCount(1)
                    .withParallelism(4)
                    .withScale(new JobScale().withPollingInterval(40).withMinExecutions(1).withMaxExecutions(5)
                        .withRules(Arrays.asList(new JobScaleRule().withName("servicebuscalingrule")
                            .withType("azure-servicebus")
                            .withMetadata(SerializerFactory.createDefaultManagementSerializerAdapter()
                                .deserialize("{\"topicName\":\"my-topic\"}", Object.class, SerializerEncoding.JSON))
                            .withIdentity(
                                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"))))))
            .withTemplate(new JobTemplate()
                .withInitContainers(Arrays.asList(new InitContainer().withImage("repo/testcontainerappsjob0:v4")
                    .withName("testinitcontainerAppsJob0").withCommand(Arrays.asList("/bin/sh"))
                    .withArgs(Arrays.asList("-c", "while true; do echo hello; sleep 10;done"))
                    .withResources(new ContainerResources().withCpu(0.5D).withMemory("1Gi"))))
                .withContainers(Arrays.asList(
                    new Container().withImage("repo/testcontainerappsjob0:v1").withName("testcontainerappsjob0"))))
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
