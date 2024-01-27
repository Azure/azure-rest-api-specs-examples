
import com.azure.resourcemanager.appplatform.models.ConfigServerGitProperty;
import com.azure.resourcemanager.appplatform.models.ConfigServerSettings;
import java.util.Arrays;

/**
 * Samples for ConfigServers Validate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ConfigServers_Validate.json
     */
    /**
     * Sample code: ConfigServers_Validate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configServersValidate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getConfigServers().validate("myResourceGroup", "myservice",
            new ConfigServerSettings().withGitProperty(
                new ConfigServerGitProperty().withUri("https://github.com/fake-user/fake-repository.git")
                    .withLabel("master").withSearchPaths(Arrays.asList("/"))),
            com.azure.core.util.Context.NONE);
    }
}
