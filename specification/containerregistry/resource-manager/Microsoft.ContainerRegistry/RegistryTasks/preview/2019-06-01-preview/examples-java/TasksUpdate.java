
import com.azure.resourcemanager.containerregistry.models.AgentProperties;
import com.azure.resourcemanager.containerregistry.models.AuthInfoUpdateParameters;
import com.azure.resourcemanager.containerregistry.models.Credentials;
import com.azure.resourcemanager.containerregistry.models.CustomRegistryCredentials;
import com.azure.resourcemanager.containerregistry.models.DockerBuildStepUpdateParameters;
import com.azure.resourcemanager.containerregistry.models.SecretObject;
import com.azure.resourcemanager.containerregistry.models.SecretObjectType;
import com.azure.resourcemanager.containerregistry.models.SourceTriggerEvent;
import com.azure.resourcemanager.containerregistry.models.SourceTriggerUpdateParameters;
import com.azure.resourcemanager.containerregistry.models.SourceUpdateParameters;
import com.azure.resourcemanager.containerregistry.models.TaskStatus;
import com.azure.resourcemanager.containerregistry.models.TaskUpdateParameters;
import com.azure.resourcemanager.containerregistry.models.TokenType;
import com.azure.resourcemanager.containerregistry.models.TriggerUpdateParameters;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Tasks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/TasksUpdate.json
     */
    /**
     * Sample code: Tasks_Update.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tasksUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTasks().update("myResourceGroup", "myRegistry",
            "myTask",
            new TaskUpdateParameters().withTags(mapOf("testkey", "fakeTokenPlaceholder")).withStatus(TaskStatus.ENABLED)
                .withAgentConfiguration(new AgentProperties().withCpu(3))
                .withStep(new DockerBuildStepUpdateParameters().withImageNames(Arrays.asList("azurerest:testtag1"))
                    .withDockerFilePath("src/DockerFile"))
                .withTrigger(
                    new TriggerUpdateParameters().withSourceTriggers(Arrays.asList(new SourceTriggerUpdateParameters()
                        .withSourceRepository(
                            new SourceUpdateParameters().withSourceControlAuthProperties(new AuthInfoUpdateParameters()
                                .withTokenType(TokenType.PAT).withToken("fakeTokenPlaceholder")))
                        .withSourceTriggerEvents(Arrays.asList(SourceTriggerEvent.COMMIT))
                        .withName("mySourceTrigger"))))
                .withCredentials(new Credentials().withCustomRegistries(mapOf("myregistry.azurecr.io",
                    new CustomRegistryCredentials()
                        .withUsername(new SecretObject().withValue("username").withType(SecretObjectType.OPAQUE))
                        .withPassword(
                            new SecretObject().withValue("https://myacbvault.vault.azure.net/secrets/password")
                                .withType(SecretObjectType.VAULTSECRET))
                        .withIdentity("[system]"))))
                .withLogTemplate("acr/tasks:{{.Run.OS}}"),
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
