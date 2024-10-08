
import com.azure.resourcemanager.containerregistry.fluent.models.TaskInner;
import com.azure.resourcemanager.containerregistry.models.AgentProperties;
import com.azure.resourcemanager.containerregistry.models.Architecture;
import com.azure.resourcemanager.containerregistry.models.Argument;
import com.azure.resourcemanager.containerregistry.models.AuthInfo;
import com.azure.resourcemanager.containerregistry.models.BaseImageTrigger;
import com.azure.resourcemanager.containerregistry.models.BaseImageTriggerType;
import com.azure.resourcemanager.containerregistry.models.DockerTaskStep;
import com.azure.resourcemanager.containerregistry.models.IdentityProperties;
import com.azure.resourcemanager.containerregistry.models.OS;
import com.azure.resourcemanager.containerregistry.models.PlatformProperties;
import com.azure.resourcemanager.containerregistry.models.ResourceIdentityType;
import com.azure.resourcemanager.containerregistry.models.SourceControlType;
import com.azure.resourcemanager.containerregistry.models.SourceProperties;
import com.azure.resourcemanager.containerregistry.models.SourceTrigger;
import com.azure.resourcemanager.containerregistry.models.SourceTriggerEvent;
import com.azure.resourcemanager.containerregistry.models.TaskStatus;
import com.azure.resourcemanager.containerregistry.models.TimerTrigger;
import com.azure.resourcemanager.containerregistry.models.TokenType;
import com.azure.resourcemanager.containerregistry.models.TriggerProperties;
import com.azure.resourcemanager.containerregistry.models.UpdateTriggerPayloadType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Tasks Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/
     * TasksCreate.json
     */
    /**
     * Sample code: Tasks_Create.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tasksCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTasks()
            .create("myResourceGroup", "myRegistry", "mytTask",
                new TaskInner().withLocation("eastus").withTags(mapOf("testkey", "fakeTokenPlaceholder"))
                    .withIdentity(new IdentityProperties().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
                    .withStatus(TaskStatus.ENABLED)
                    .withPlatform(new PlatformProperties().withOs(OS.LINUX).withArchitecture(Architecture.AMD64))
                    .withAgentConfiguration(new AgentProperties().withCpu(2))
                    .withStep(new DockerTaskStep()
                        .withContextPath("src").withImageNames(Arrays.asList("azurerest:testtag")).withIsPushEnabled(
                            true)
                        .withNoCache(false).withDockerFilePath("src/DockerFile")
                        .withArguments(Arrays.asList(
                            new Argument().withName("mytestargument").withValue("mytestvalue").withIsSecret(false),
                            new Argument().withName("mysecrettestargument").withValue("mysecrettestvalue")
                                .withIsSecret(true))))
                    .withTrigger(
                        new TriggerProperties()
                            .withTimerTriggers(
                                Arrays.asList(new TimerTrigger().withSchedule("30 9 * * 1-5")
                                    .withName("myTimerTrigger")))
                            .withSourceTriggers(Arrays.asList(new SourceTrigger()
                                .withSourceRepository(
                                    new SourceProperties().withSourceControlType(SourceControlType.GITHUB)
                                        .withRepositoryUrl("https://github.com/Azure/azure-rest-api-specs")
                                        .withBranch("master")
                                        .withSourceControlAuthProperties(new AuthInfo().withTokenType(TokenType.PAT)
                                            .withToken("fakeTokenPlaceholder")))
                                .withSourceTriggerEvents(Arrays.asList(SourceTriggerEvent.COMMIT))
                                .withName("mySourceTrigger")))
                            .withBaseImageTrigger(
                                new BaseImageTrigger().withBaseImageTriggerType(BaseImageTriggerType.RUNTIME)
                                    .withUpdateTriggerEndpoint("https://user:pass@mycicd.webhook.com?token=foo")
                                    .withUpdateTriggerPayloadType(UpdateTriggerPayloadType.TOKEN)
                                    .withName("myBaseImageTrigger")))
                    .withLogTemplate("acr/tasks:{{.Run.OS}}").withIsSystemTask(false),
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
