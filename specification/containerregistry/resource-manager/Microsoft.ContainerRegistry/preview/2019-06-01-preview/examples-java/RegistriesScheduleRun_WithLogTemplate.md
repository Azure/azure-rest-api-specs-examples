Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.models.AgentProperties;
import com.azure.resourcemanager.containerregistry.models.Architecture;
import com.azure.resourcemanager.containerregistry.models.Argument;
import com.azure.resourcemanager.containerregistry.models.DockerBuildRequest;
import com.azure.resourcemanager.containerregistry.models.OS;
import com.azure.resourcemanager.containerregistry.models.PlatformProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Registries ScheduleRun. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RegistriesScheduleRun_WithLogTemplate.json
     */
    /**
     * Sample code: Registries_ScheduleRun_WithLogTemplate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registriesScheduleRunWithLogTemplate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .scheduleRun(
                "myResourceGroup",
                "myRegistry",
                new DockerBuildRequest()
                    .withIsArchiveEnabled(true)
                    .withLogTemplate("acr/tasks:{{.Run.OS}}")
                    .withImageNames(Arrays.asList("azurerest:testtag"))
                    .withIsPushEnabled(true)
                    .withNoCache(true)
                    .withDockerFilePath("DockerFile")
                    .withArguments(
                        Arrays
                            .asList(
                                new Argument().withName("mytestargument").withValue("mytestvalue").withIsSecret(false),
                                new Argument()
                                    .withName("mysecrettestargument")
                                    .withValue("mysecrettestvalue")
                                    .withIsSecret(true)))
                    .withPlatform(new PlatformProperties().withOs(OS.LINUX).withArchitecture(Architecture.AMD64))
                    .withAgentConfiguration(new AgentProperties().withCpu(2))
                    .withSourceLocation(
                        "https://myaccount.blob.core.windows.net/sascontainer/source.zip?sv=2015-04-05&st=2015-04-29T22%3A18%3A26Z&se=2015-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=Z%2FRHIX5Xcg0Mq2rqI3OlWTjEg2tYkboXr1P9ZUXDtkk%3D"),
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
