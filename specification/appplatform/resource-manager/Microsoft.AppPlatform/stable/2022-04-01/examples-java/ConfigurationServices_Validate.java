import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.models.ConfigurationServiceGitProperty;
import com.azure.resourcemanager.appplatform.models.ConfigurationServiceGitRepository;
import com.azure.resourcemanager.appplatform.models.ConfigurationServiceSettings;
import java.util.Arrays;

/** Samples for ConfigurationServices Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ConfigurationServices_Validate.json
     */
    /**
     * Sample code: ConfigurationServices_Validate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configurationServicesValidate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getConfigurationServices()
            .validate(
                "myResourceGroup",
                "myservice",
                "default",
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
                                            .withLabel("master")))),
                Context.NONE);
    }
}
