Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.models.Architecture;
import com.azure.resourcemanager.containerregistry.models.Credentials;
import com.azure.resourcemanager.containerregistry.models.EncodedTaskRunRequest;
import com.azure.resourcemanager.containerregistry.models.OS;
import com.azure.resourcemanager.containerregistry.models.PlatformProperties;
import com.azure.resourcemanager.containerregistry.models.TaskRunUpdateParameters;
import java.util.Arrays;

/** Samples for TaskRuns Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsUpdate.json
     */
    /**
     * Sample code: TaskRuns_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void taskRunsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getTaskRuns()
            .update(
                "myResourceGroup",
                "myRegistry",
                "myRun",
                new TaskRunUpdateParameters()
                    .withRunRequest(
                        new EncodedTaskRunRequest()
                            .withIsArchiveEnabled(true)
                            .withEncodedTaskContent("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K")
                            .withEncodedValuesContent("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg==")
                            .withValues(Arrays.asList())
                            .withPlatform(
                                new PlatformProperties().withOs(OS.LINUX).withArchitecture(Architecture.AMD64))
                            .withCredentials(new Credentials()))
                    .withForceUpdateTag("test"),
                Context.NONE);
    }
}
```
