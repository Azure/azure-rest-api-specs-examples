import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.ConfigurationServiceResourceInner;
import com.azure.resourcemanager.appplatform.models.ConfigurationServiceGitProperty;
import com.azure.resourcemanager.appplatform.models.ConfigurationServiceGitRepository;
import com.azure.resourcemanager.appplatform.models.ConfigurationServiceProperties;
import com.azure.resourcemanager.appplatform.models.ConfigurationServiceSettings;
import java.util.Arrays;

/** Samples for ConfigurationServices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ConfigurationServices_CreateOrUpdate.json
     */
    /**
     * Sample code: ConfigurationServices_CreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configurationServicesCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getConfigurationServices()
            .createOrUpdate(
                "myResourceGroup",
                "myservice",
                "default",
                new ConfigurationServiceResourceInner()
                    .withProperties(
                        new ConfigurationServiceProperties()
                            .withSettings(
                                new ConfigurationServiceSettings()
                                    .withGitProperty(
                                        new ConfigurationServiceGitProperty()
                                            .withRepositories(
                                                Arrays
                                                    .asList(
                                                        new ConfigurationServiceGitRepository()
                                                            .withName("fake")
                                                            .withPatterns(Arrays.asList("app/dev"))
                                                            .withUri("https://github.com/fake-user/fake-repository")
                                                            .withLabel("master")))))),
                Context.NONE);
    }
}
