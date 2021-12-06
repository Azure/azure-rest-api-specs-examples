Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
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
import com.azure.resourcemanager.containerregistry.models.UserIdentityProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Tasks Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/ManagedIdentity/TasksCreate_WithUserIdentities.json
     */
    /**
     * Sample code: Tasks_Create_WithUserIdentities.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tasksCreateWithUserIdentities(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getTasks()
            .create(
                "myResourceGroup",
                "myRegistry",
                "mytTask",
                new TaskInner()
                    .withLocation("eastus")
                    .withTags(mapOf("testkey", "value"))
                    .withIdentity(
                        new IdentityProperties()
                            .withType(ResourceIdentityType.USER_ASSIGNED)
                            .withUserAssignedIdentities(
                                mapOf(
                                    "/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1",
                                    new UserIdentityProperties(),
                                    "/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2",
                                    new UserIdentityProperties())))
                    .withStatus(TaskStatus.ENABLED)
                    .withPlatform(new PlatformProperties().withOs(OS.LINUX).withArchitecture(Architecture.AMD64))
                    .withAgentConfiguration(new AgentProperties().withCpu(2))
                    .withStep(
                        new DockerTaskStep()
                            .withContextPath("src")
                            .withImageNames(Arrays.asList("azurerest:testtag"))
                            .withIsPushEnabled(true)
                            .withNoCache(false)
                            .withDockerFilePath("src/DockerFile")
                            .withArguments(
                                Arrays
                                    .asList(
                                        new Argument()
                                            .withName("mytestargument")
                                            .withValue("mytestvalue")
                                            .withIsSecret(false),
                                        new Argument()
                                            .withName("mysecrettestargument")
                                            .withValue("mysecrettestvalue")
                                            .withIsSecret(true))))
                    .withTrigger(
                        new TriggerProperties()
                            .withTimerTriggers(
                                Arrays
                                    .asList(new TimerTrigger().withSchedule("30 9 * * 1-5").withName("myTimerTrigger")))
                            .withSourceTriggers(
                                Arrays
                                    .asList(
                                        new SourceTrigger()
                                            .withSourceRepository(
                                                new SourceProperties()
                                                    .withSourceControlType(SourceControlType.GITHUB)
                                                    .withRepositoryUrl("https://github.com/Azure/azure-rest-api-specs")
                                                    .withBranch("master")
                                                    .withSourceControlAuthProperties(
                                                        new AuthInfo().withTokenType(TokenType.PAT).withToken("xxxxx")))
                                            .withSourceTriggerEvents(Arrays.asList(SourceTriggerEvent.COMMIT))
                                            .withName("mySourceTrigger")))
                            .withBaseImageTrigger(
                                new BaseImageTrigger()
                                    .withBaseImageTriggerType(BaseImageTriggerType.RUNTIME)
                                    .withUpdateTriggerEndpoint("https://user:pass@mycicd.webhook.com?token=foo")
                                    .withUpdateTriggerPayloadType(UpdateTriggerPayloadType.DEFAULT)
                                    .withName("myBaseImageTrigger")))
                    .withIsSystemTask(false),
                Context.NONE);
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
```
