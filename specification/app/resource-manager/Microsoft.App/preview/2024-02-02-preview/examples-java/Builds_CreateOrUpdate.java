
import com.azure.resourcemanager.appcontainers.models.BuildConfiguration;
import com.azure.resourcemanager.appcontainers.models.ContainerRegistryWithCustomImage;
import com.azure.resourcemanager.appcontainers.models.EnvironmentVariable;
import com.azure.resourcemanager.appcontainers.models.HttpGet;
import com.azure.resourcemanager.appcontainers.models.PreBuildStep;
import java.util.Arrays;

/**
 * Samples for Builds CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/Builds_CreateOrUpdate.json
     */
    /**
     * Sample code: Builds_CreateOrUpdate_WithConfig.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        buildsCreateOrUpdateWithConfig(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.builds().define("testBuild-123456789az").withExistingBuilder("rg", "testBuilder")
            .withDestinationContainerRegistry(new ContainerRegistryWithCustomImage().withServer("test.azurecr.io")
                .withImage("test.azurecr.io/repo:tag"))
            .withConfiguration(
                new BuildConfiguration().withBaseOs("DebianBullseye").withPlatform("dotnetcore")
                    .withPlatformVersion("7.0")
                    .withEnvironmentVariables(
                        Arrays.asList(new EnvironmentVariable().withName("foo1").withValue("bar1"),
                            new EnvironmentVariable().withName("foo2").withValue("bar2")))
                    .withPreBuildSteps(Arrays.asList(
                        new PreBuildStep().withDescription("First pre build step.")
                            .withScripts(Arrays.asList("echo 'hello'", "echo 'world'"))
                            .withHttpGet(new HttpGet().withUrl("https://microsoft.com").withFileName("output.txt")
                                .withHeaders(Arrays.asList("foo", "bar"))),
                        new PreBuildStep().withDescription("Second pre build step.")
                            .withScripts(Arrays.asList("echo 'hello'", "echo 'again'"))
                            .withHttpGet(new HttpGet().withUrl("https://microsoft.com").withFileName("output.txt")
                                .withHeaders(Arrays.asList("foo"))))))
            .create();
    }
}
